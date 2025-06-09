package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/middleware"
	authhandlerv1 "starter-go-gin/modules/auth/v1/handler"
	productHandlerv1 "starter-go-gin/modules/product/v1/handler"
	utilscreatorhandlerv1 "starter-go-gin/modules/utils/v1/handler"
	"starter-go-gin/response"
)

// DeprecatedAPI is a handler for deprecated APIs
func DeprecatedAPI(c *gin.Context) {
	c.JSON(http.StatusForbidden, response.ErrorAPIResponse(http.StatusForbidden, "this version of api is deprecated. please use another version."))
	c.Abort()
}

// DefaultHTTPHandler is a handler for default APIs
func DefaultHTTPHandler(cfg config.Config, router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.ErrorAPIResponse(http.StatusNotFound, "invalid route"))
		c.Abort()
	})
}

// AuthHTTPHandler is a handler for auth APIs
func AuthHTTPHandler(cfg config.Config, router *gin.Engine,
	hnd *authhandlerv1.AuthHandler, cache interfaces.Cacheable,
) {
	v1 := router.Group("/v1")
	{
		// user app
		v1.POST("/user/login", hnd.Login)
		v1.POST("/user/register", hnd.Register)

		// admin app
		v1.POST("/admin/login", hnd.LoginAdmin)
	}
}

// ProductCreatorHTTPHandler is a handler for product APIs
func ProductCreatorHTTPHandler(cfg config.Config, router *gin.Engine,
	hnd *productHandlerv1.ProductCreatorHandler, cache interfaces.Cacheable,
) {

	v1Admin := router.Group("/v1/product", middleware.Admin(cfg))
	{
		v1Admin.POST("", hnd.CreateProduct)
	}
}

// NotificationFinderHTTPHandler is a handler for notification APIs
func NotificationFinderHTTPHandler(
	cfg config.Config,
	router *gin.Engine,
	cache interfaces.Cacheable,
	hnd *productHandlerv1.ProductCreatorHandler) {

	v1Admin := router.Group("/v1/product", middleware.Admin(cfg))
	{
		v1Admin.GET("/notifications", hnd.CreateProduct)
	}
}

// UtilsCreatorHTTPHandler is a handler for utils APIs
func UtilsCreatorHTTPHandler(cfg config.Config, router *gin.Engine, hnd *utilscreatorhandlerv1.UtilsCreatorHandler) {
	v1 := router.Group("/v1")
	{
		v1.POST("/utils/upload-file", hnd.UploadFile)
	}
}
