package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opencensus.io/trace"

	"starter-go-gin/common/constant"
	"starter-go-gin/common/errors"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/tools"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/middleware"
	"starter-go-gin/modules/auth/v1/service"
	service2 "starter-go-gin/modules/user/v1/service"
	"starter-go-gin/resource"
	"starter-go-gin/response"
)

// AuthHandler is a handler for auth
type AuthHandler struct {
	cfg               *config.Config
	authUseCase       service.AuthUseCase
	userFinderService service2.UserFinderUseCase
	cache             interfaces.Cacheable
}

// NewAuthHandler creates a new AuthHandler
func NewAuthHandler(
	cfg *config.Config,
	authUseCase service.AuthUseCase,
	userFinderService service2.UserFinderUseCase,
	cache interfaces.Cacheable,
) *AuthHandler {
	return &AuthHandler{
		cfg:               cfg,
		authUseCase:       authUseCase,
		userFinderService: userFinderService,
		cache:             cache,
	}
}

// Login is a handler for login user
func (ah *AuthHandler) Login(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "Login")
	defer span.End()
	requestId := middleware.GetRequestID(ctx)

	var request resource.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, requestId, parseError[0]))
		c.Abort()
		return
	}

	var res *entity.User
	var err error

	res, err = ah.authUseCase.AuthValidate(ctx, request.Username, request.Password, constant.EmployeeRoleName)

	if err != nil && err.Error() == errors.ErrRecordNotFound.Error().Error() {
		err = errors.ErrLoginNotFound.Error()
	}

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, requestId, parseError.Message))
		c.Abort()
		return
	}

	// generate token
	token, err := ah.authUseCase.GenerateAccessToken(ctx, res)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, requestId, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", resource.NewLoginResponse(token.Token)))
}

// Login is a handler for login admin
func (ah *AuthHandler) LoginAdmin(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "Login")
	defer span.End()
	requestId := middleware.GetRequestID(ctx)

	var request resource.LoginRequestAdmin

	if err := c.ShouldBind(&request); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, requestId, parseError[0]))
		c.Abort()
		return
	}

	var res *entity.User
	var err error

	res, err = ah.authUseCase.AuthValidateAdmin(ctx, request.Username, request.Password)

	if err != nil && err.Error() == errors.ErrRecordNotFound.Error().Error() {
		err = errors.ErrLoginNotFound.Error()
	}

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, requestId, parseError.Message))
		c.Abort()
		return
	}

	// generate token
	token, err := ah.authUseCase.GenerateAccessTokenAdmin(ctx, res)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, requestId, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", resource.NewLoginResponse(token.Token)))
}

// Logout is a handler for logout
func (ah *AuthHandler) Logout(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "Logout")
	defer span.End()
	fmt.Println("ctx:", ctx)
	requestId := middleware.GetRequestID(ctx)

	var req *resource.LogoutRequest
	if err := c.ShouldBind(&req); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, requestId, parseError[0]))
		return
	}

	deviceID := tools.EscapeSpecial(req.DeviceID)
	err := ah.authUseCase.Logout(ctx, uuid.MustParse(c.GetHeader(constant.RequestHeaderUserID)), deviceID, uuid.MustParse(c.GetHeader(constant.RequestHeaderMerchantID)))
	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, requestId, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, requestId, "success", nil))
}
