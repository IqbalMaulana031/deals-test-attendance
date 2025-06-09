package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/product/v1/repository"
)

// ProductFinder is a service for product
type ProductFinder struct {
	pfg         config.Config
	productRepo repository.ProductRepositoryUseCase
}

// ProductFinderUseCase is a usecase for product
type ProductFinderUseCase interface {
	// GetProduct finds all product by filter
	GetProduct(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Product, int64, error)
	// GetProductByID finds an product by id
	GetProductByID(ctx context.Context, id uuid.UUID) (*entity.Product, error)
}

// NewProductFinder creates a new ProductFinder
func NewProductFinder(
	pfg config.Config,
	productRepo repository.ProductRepositoryUseCase,
) *ProductFinder {
	return &ProductFinder{
		pfg:         pfg,
		productRepo: productRepo,
	}
}

// GetProductByID gets a admin product by ID
func (pf *ProductFinder) GetProductByID(ctx context.Context, id uuid.UUID) (*entity.Product, error) {
	product, err := pf.productRepo.GetProductByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	if product == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return product, nil
}

// GetProduct finds all product by filter
func (pf *ProductFinder) GetProduct(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Product, int64, error) {
	product, total, err := pf.productRepo.GetProduct(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}

	return product, total, nil
}
