package handler

import (
	"net/http"
	"starter-go-gin/common/errors"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/product/v1/service"
	"starter-go-gin/resource"
	"starter-go-gin/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opencensus.io/trace"
)

// ProductCreatorHandler is a handler for product finder
type ProductCreatorHandler struct {
	cfg            config.Config
	productCreator service.ProductCreatorUseCase
	cache          interfaces.Cacheable
}

// NewProductCreatorHandler is a constructor for ProductCreatorHandler
func NewProductCreatorHandler(
	cfg config.Config,
	productCreator service.ProductCreatorUseCase,
	cache interfaces.Cacheable,
) *ProductCreatorHandler {
	return &ProductCreatorHandler{
		cfg:            cfg,
		productCreator: productCreator,
		cache:          cache,
	}
}

// CreateProduct is a handler for create product
func (pch *ProductCreatorHandler) CreateProduct(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "handler.productCreator.Create")
	defer span.End()

	var req resource.CreateProductRequest

	if err := c.ShouldBind(&req); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, parseError[0]))
		c.Abort()
		return
	}

	product := entity.NewProduct(
		uuid.New(),
		req.ProductName,
		req.Stock,
		req.Price,
		"system",
	)
	err := pch.productCreator.CreateProduct(ctx, product)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", nil))
}
