package service

import (
	"context"

	"starter-go-gin/common/errors"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/product/v1/repository"
)

// ProductCreator is a struct that contains all the dependencies for the User creator
type ProductCreator struct {
	cfg         config.Config
	productRepo repository.ProductRepositoryUseCase
}

// ProductCreatorUseCase is a use case for the User creator
type ProductCreatorUseCase interface {
	// CreateProduct create product
	CreateProduct(ctx context.Context, product *entity.Product) error
}

// NewProductCreator is a constructor for the User creator
func NewProductCreator(
	cfg config.Config,
	productRepo repository.ProductRepositoryUseCase,
) *ProductCreator {
	return &ProductCreator{
		cfg:         cfg,
		productRepo: productRepo,
	}
}

// DeleteRole deletes role
func (ud *ProductCreator) CreateProduct(ctx context.Context, product *entity.Product) error {
	if err := ud.productRepo.CreateProduct(ctx, product); err != nil {
		return errors.ErrInternalServerError.Error()
	}
	return nil
}
