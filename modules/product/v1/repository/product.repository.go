package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	commonCache "starter-go-gin/common/cache"
	"starter-go-gin/common/constant"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/logger"
	"starter-go-gin/common/tools"
	"starter-go-gin/config"
	"starter-go-gin/entity"
)

// ProductRepository is a repository for auth
type ProductRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
	// The code below is commented out because synchronization with Firestore is no longer needed
	// firestore interfaces.FirestoreUseCase
	cfg config.Config
}

// ProductRepositoryUseCase is a repository for auth
type ProductRepositoryUseCase interface {
	// GetProduct finds all product by filter
	GetProduct(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Product, int64, error)
	// GetProductByID finds an product by id
	GetProductByID(ctx context.Context, id uuid.UUID) (*entity.Product, error)
	// CreateProduct create product
	CreateProduct(ctx context.Context, product *entity.Product) error
	// UpdateProduct update product
	UpdateProduct(ctx context.Context, product *entity.Product, id uuid.UUID) error
	// DeleteProduct update stock
	DeleteProduct(ctx context.Context, id uuid.UUID) error
}

// NewProductRepository returns a auth repository
func NewProductRepository(
	db *gorm.DB,
	cache interfaces.Cacheable,
	cfg config.Config,
) *ProductRepository {
	return &ProductRepository{db, cache, cfg}
}

// GetProduct is a function to get all product
func (ur *ProductRepository) GetProduct(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Product, int64, error) {
	var product []*entity.Product
	var total int64
	var gormDB = ur.db.
		WithContext(ctx).
		Model(&entity.Product{}).
		Where("product.products.stock > 0")

	gormDB.Count(&total)

	gormDB = gormDB.Limit(limit).
		Offset(offset * limit)

	if query != "" {
		gormDB = gormDB.
			Where("product_name ILIKE ?", "%"+query+"%")
	}

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", tools.EscapeSpecial(sort), tools.EscapeSpecial(order)))

	if err := gormDB.Find(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[ProductRepository-GetProduct] error when looking up all product")
	}

	return product, total, nil
}

// GetProductByID finds an product by id
func (ur *ProductRepository) GetProductByID(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	result := &entity.Product{}

	if err := ur.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logger.Error(ctx, errors.Wrap(err, "[ProductRepository-GetProductByID] id not found"))
		return nil, errors.Wrap(err, "[ProductRepository-GetProductByID] id not found")
	}

	return result, nil
}

// CreateProduct create product
func (ur *ProductRepository) CreateProduct(ctx context.Context, product *entity.Product) error {
	if err := ur.db.
		WithContext(ctx).
		Model(&entity.Product{}).
		Create(product).
		Error; err != nil {
		return errors.Wrap(err, "[ProductRepository-CreateProduct] error while creating prduct")
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.ProductFindByID, product.ID)); err != nil {
		return err
	}

	return nil
}

// UpdateProduct is a function to update product
func (ur *ProductRepository) UpdateProduct(ctx context.Context, product *entity.Product, id uuid.UUID) error {
	oldTime := product.UpdatedAt
	product.UpdatedAt = time.Now()
	if err := ur.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.Product)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, product.ID).Error; err != nil {
				logger.Error(ctx, err)
				return err
			}
			if err := tx.Model(&entity.Product{}).
				Where(`id`, product.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(product)).Error; err != nil {
				logger.Error(ctx, err)
				return err
			}
			return nil
		}); err != nil {
		product.UpdatedAt = oldTime
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.ProductFindByID, product.ID)); err != nil {
		return err
	}

	return nil
}

// DeleteProduct is a function to delete product
func (ur *ProductRepository) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if err := ur.db.WithContext(ctx).
		Model(&entity.Product{}).
		Where(`id = ?`, id).
		Delete(&entity.Product{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "[ProductRepository-DeleteProduct] error when updating product data")
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.ProductFindByID, "*")); err != nil {
		return err
	}

	return nil
}
