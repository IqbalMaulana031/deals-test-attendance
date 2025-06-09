package handler

import (
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
	"starter-go-gin/modules/auth/v1/service"
	service2 "starter-go-gin/modules/user/v1/service"
	"starter-go-gin/resource"
	"starter-go-gin/response"
	"starter-go-gin/utils"
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

	var request resource.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, parseError[0]))
		c.Abort()
		return
	}

	var res *entity.User
	var err error

	res, err = ah.authUseCase.AuthValidate(ctx, request.UsernameOrEmail, request.Password, constant.RoleUser)

	if err != nil && err.Error() == errors.ErrRecordNotFound.Error().Error() {
		err = errors.ErrLoginNotFound.Error()
	}

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	// generate token
	token, err := ah.authUseCase.GenerateAccessToken(ctx, res)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewLoginResponse(token.Token)))
}

// Login is a handler for login admin
func (ah *AuthHandler) LoginAdmin(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "Login")
	defer span.End()

	var request resource.LoginRequestAdmin

	if err := c.ShouldBind(&request); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, parseError[0]))
		c.Abort()
		return
	}

	var res *entity.User
	var err error

	res, err = ah.authUseCase.AuthValidateAdmin(ctx, request.Email, request.Password)

	if err != nil && err.Error() == errors.ErrRecordNotFound.Error().Error() {
		err = errors.ErrLoginNotFound.Error()
	}

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	// generate token
	token, err := ah.authUseCase.GenerateAccessTokenCMS(ctx, res)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", resource.NewLoginResponse(token.Token)))
}

// Register is a handler for register user
func (ah *AuthHandler) Register(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "Register")
	defer span.End()

	var req resource.RegisterRequest

	if err := c.ShouldBind(&req); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, parseError[0]))
		c.Abort()
		return
	}

	emailDomain, err := utils.GetDomainSubstring(req.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, err.Error()))
		c.Abort()
		return
	}

	if !utils.FindString(constant.ListPopularEmail, emailDomain) {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, "Invalid email domain"))
		c.Abort()
		return
	}

	// ===============================CHECK IF USER EXIST IN MERCHANT (merchant_id is not null) ===============================
	userEmail, _ := ah.userFinderService.GetUserByEmail(ctx, req.Email, constant.RoleUser)

	if userEmail != nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, errors.ErrEmailAlreadyExist.Message))
		c.Abort()
		return
	}
	// ===============================================================================================

	// check if password and duplicate password is same
	if req.Password != req.DuplicatePassword {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, "Different password and duplicate password"))
		c.Abort()
		return
	}

	user := entity.NewUser(
		uuid.New(),
		req.Username,
		req.Email,
		req.Password,
		req.Gender,
		req.Username,
	)

	// find role merchant owner
	role, err := ah.userFinderService.GetRoleByName(ctx, constant.RoleUser)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	if role == nil {
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, "role merchant owner not found"))
		c.Abort()
		return
	}

	userRole := entity.NewUserRole(
		uuid.New(),
		user.ID,
		role.ID,
		constant.CreatedBySystem,
	)

	_, err = ah.authUseCase.Register(ctx, user, userRole)

	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", nil))
}

// Logout is a handler for logout
func (ah *AuthHandler) Logout(c *gin.Context) {
	ctx, span := trace.StartSpan(c.Request.Context(), "Logout")
	defer span.End()

	var req *resource.LogoutRequest
	if err := c.ShouldBind(&req); err != nil {
		parseError := errors.ParseErrorValidation(err)
		c.JSON(http.StatusBadRequest, response.ErrorAPIResponse(http.StatusBadRequest, parseError[0]))
		return
	}

	deviceID := tools.EscapeSpecial(req.DeviceID)
	err := ah.authUseCase.Logout(ctx, uuid.MustParse(c.GetHeader(constant.RequestHeaderUserID)), deviceID, uuid.MustParse(c.GetHeader(constant.RequestHeaderMerchantID)))
	if err != nil {
		parseError := errors.ParseError(err)
		c.JSON(parseError.Code, response.ErrorAPIResponse(parseError.Code, parseError.Message))
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, response.SuccessAPIResponseList(http.StatusOK, "success", nil))
}
