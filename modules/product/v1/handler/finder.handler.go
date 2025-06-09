package handler

import (
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/modules/product/v1/service"
)

// ProductFinderHandler is a handler for product finder
type ProductFinderHandler struct {
	productFinder service.ProductFinderUseCase
	cfg           config.Config
	cache         interfaces.Cacheable
	excelize      interfaces.ExcelizeUseCase
}

// NewProductFinderHandler is a constructor for ProductFinderHandler
func NewProductFinderHandler(
	productFinder service.ProductFinderUseCase,
	cache interfaces.Cacheable,
	cfg config.Config,
	gotenberg interfaces.GotenbergUseCase,
	excelize interfaces.ExcelizeUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *ProductFinderHandler {
	return &ProductFinderHandler{
		productFinder: productFinder,
		cache:         cache,
		excelize:      excelize,
	}
}
