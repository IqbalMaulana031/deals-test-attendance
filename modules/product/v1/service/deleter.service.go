package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/config"
	"starter-go-gin/modules/product/v1/repository"
)

// ProductDeleter is a service for product
type ProductDeleter struct {
	cfg         config.Config
	productRepo repository.ProductRepositoryUseCase
}

// ProductDeleterUseCase is a use case for product
type ProductDeleterUseCase interface {
	// DeleteProduct update stock
	DeleteProduct(ctx context.Context, ID uuid.UUID) error
}

// NewProductDeleter creates a new ProductDeleter
func NewProductDeleter(
	cfg config.Config,
	productRepo repository.ProductRepositoryUseCase,
) *ProductDeleter {
	return &ProductDeleter{
		cfg:         cfg,
		productRepo: productRepo,
	}
}

// DeleteProduct deletes product
func (ud *ProductDeleter) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if err := ud.productRepo.DeleteProduct(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
