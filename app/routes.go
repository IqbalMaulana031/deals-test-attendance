package app

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/middleware"
	attendanceHandlerv1 "starter-go-gin/modules/attendance/v1/handler"
	authhandlerv1 "starter-go-gin/modules/auth/v1/handler"
	masterHandlerv1 "starter-go-gin/modules/master/v1/handler"
	utilscreatorhandlerv1 "starter-go-gin/modules/utils/v1/handler"
	"starter-go-gin/response"
)

// DeprecatedAPI is a handler for deprecated APIs
func DeprecatedAPI(c *gin.Context) {
	c.JSON(http.StatusForbidden, response.ErrorAPIResponseWithoutReqID(http.StatusForbidden, "this version of api is deprecated. please use another version."))
	c.Abort()
}

// DefaultHTTPHandler is a handler for default APIs
func DefaultHTTPHandler(cfg config.Config, router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.ErrorAPIResponseWithoutReqID(http.StatusNotFound, "invalid route"))
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

		// admin app
		v1.POST("/admin/login", hnd.LoginAdmin)
	}
}

// MasterCreatorHTTPHandler is a handler for product APIs
func MasterCreatorHTTPHandler(cfg config.Config, router *gin.Engine,
	hnd *masterHandlerv1.MasterCreatorHandler, cache interfaces.Cacheable,
) {

	// v1Admin := router.Group("/v1/product", middleware.Admin(cfg))
	// {
	// 	v1Admin.POST("", hnd.CreateMaster)
	// }
}

// MasterFinderHTTPHandler is a handler for notification APIs
func MasterFinderHTTPHandler(
	cfg config.Config,
	router *gin.Engine,
	cache interfaces.Cacheable,
	hnd *masterHandlerv1.MasterFinderHandler) {

	v1 := router.Group("/v1/shift")
	{
		v1.GET("", hnd.GetShift)
		v1.GET("/:id", hnd.GetShiftByID)
		v1.GET("/:id/details", hnd.GetShiftAndDetailsByID)
	}
}

// AttendanceCreatorHTTPHandler is a handler for attendance APIs
func AttendanceCreatorHTTPHandler(cfg config.Config, router *gin.Engine,
	hnd *attendanceHandlerv1.AttendanceCreatorHandler, cache interfaces.Cacheable) {

	v1 := router.Group("/v1/attendance", middleware.Auth(cfg, cache))
	{
		v1.POST("/record-time", hnd.RechordTime)
	}
}

// AttendanceFinderHTTPHandler is a handler for attendance APIs
func AttendanceFinderHTTPHandler(cfg config.Config, router *gin.Engine,
	hnd *attendanceHandlerv1.AttendanceFinderHandler, cache interfaces.Cacheable) {

	v1 := router.Group("/v1/attendance", middleware.Auth(cfg, cache))
	{
		v1.GET("", hnd.GetAttendance)
		v1.GET("/:id", hnd.GetAttendanceByID)
	}

	v1Admin := router.Group("/v1/admin/attendance", middleware.Admin(cfg))
	{
		v1Admin.GET("", hnd.GetAttendanceByUserID)
		v1Admin.GET("/:id", hnd.GetAttendanceByID)
	}
}

// UtilsCreatorHTTPHandler is a handler for utils APIs
func UtilsCreatorHTTPHandler(cfg config.Config, router *gin.Engine, hnd *utilscreatorhandlerv1.UtilsCreatorHandler) {
	v1 := router.Group("/v1")
	{
		v1.POST("/utils/upload-file", hnd.UploadFile)
	}
}
