package handler

import (
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/modules/product/v1/service"
)

// ProductDeleterHandler is a handler for product finder
type ProductDeleterHandler struct {
	cfg            config.Config
	productDeleter service.ProductDeleterUseCase
	productFinder  service.ProductFinderUseCase
	cache          interfaces.Cacheable
}

// NewProductDeleterHandler is a constructor for ProductDeleterHandler
func NewProductDeleterHandler(
	cfg config.Config,
	productDeleter service.ProductDeleterUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
	productFinder service.ProductFinderUseCase,
	cacheable interfaces.Cacheable,
) *ProductDeleterHandler {
	return &ProductDeleterHandler{
		cfg:            cfg,
		productDeleter: productDeleter,
		productFinder:  productFinder,
		cache:          cacheable,
	}
}
