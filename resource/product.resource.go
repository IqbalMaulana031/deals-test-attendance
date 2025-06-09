package resource

import (
	"starter-go-gin/entity"

	"github.com/google/uuid"
)

// CreateProductRequest is a struct for request create product
type CreateProductRequest struct {
	ProductName string `json:"product_name"`
	Stock       int64  `json:"stock"`
	Price       int64  `json:"price"`
}

// ProductResponse is a struct for response create product
type ProductResponse struct {
	ID          uuid.UUID `json:"id"`
	ProductName string    `json:"product_name"`
	Stock       int64     `json:"stock"`
	Price       int64     `json:"price"`
}

// NewProductResponse is a constructor for Product
func NewProductCreateResponse(product *entity.Product) *ProductResponse {
	return &ProductResponse{
		ID:          product.ID,
		ProductName: product.ProductName,
		Stock:       product.Stock,
		Price:       product.Price,
	}
}
