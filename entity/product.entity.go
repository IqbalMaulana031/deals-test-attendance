package entity

import (
	"github.com/google/uuid"
)

const (
	productTableName = "product.products"
)

// Product defines table product
type Product struct {
	ID          uuid.UUID `json:"id"`
	ProductName string    `json:"product_ProductName"`
	Stock       int64     `json:"stock"`
	Price       int64     `json:"price"`
	Auditable
}

// TableName specifies table ProductName
func (model *Product) TableName() string {
	return productTableName
}

// NewProduct creates new product entity
func NewProduct(
	id uuid.UUID,
	ProductName string,
	stock int64,
	price int64,
	createdBy string,
) *Product {
	return &Product{
		ID:          id,
		ProductName: ProductName,
		Stock:       stock,
		Price:       stock,
		Auditable:   NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *Product) MapUpdateFrom(from *Product) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"product_ProductName": model.ProductName,
			"stock":               model.Stock,
			"price":               model.Price,
		}
	}

	mapped := make(map[string]interface{})

	if model.ProductName != from.ProductName {
		mapped["product_ProductName"] = from.ProductName
	}

	if model.Stock != from.Stock {
		mapped["price"] = from.Stock
	}

	if model.Price != from.Price {
		mapped["price"] = from.Price
	}

	mapped["updated_at"] = from.UpdatedAt

	return &mapped
}
