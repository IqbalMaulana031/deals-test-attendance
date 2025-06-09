package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/product/v1/repository"
)

// ProdcutUpdater is a struct that contains the dependencies of ProdcutUpdater
type ProdcutUpdater struct {
	cfg         config.Config
	productRepo repository.ProductRepositoryUseCase
}

// ProductUpdaterUseCase is a struct that contains the dependencies of ProductUpdaterUseCase
type ProductUpdaterUseCase interface {
	// Update a product
	// UpdateProduct updates a product
	UpdateProduct(ctx context.Context, product *entity.Product, ID uuid.UUID) error
}

// NewProdcutUpdater is a function that creates a new ProdcutUpdater
func NewProdcutUpdater(
	cfg config.Config,
	productRepo repository.ProductRepositoryUseCase,
) *ProdcutUpdater {
	return &ProdcutUpdater{
		cfg:         cfg,
		productRepo: productRepo,
	}
}

// UpdateProduct updates a product
func (uu *ProdcutUpdater) UpdateProduct(ctx context.Context, product *entity.Product, ID uuid.UUID) error {
	err := uu.productRepo.UpdateProduct(ctx, product, ID)

	if err != nil {
		return errors.ErrInternalServerError.Error()
	}

	if product == nil {
		return errors.ErrRecordNotFound.Error()
	}

	return nil
}
