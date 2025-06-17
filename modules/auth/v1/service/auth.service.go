package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	commonCache "starter-go-gin/common/cache"
	"starter-go-gin/common/constant"
	"starter-go-gin/common/errors"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/logger"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/auth/v1/repository"
	"starter-go-gin/modules/user/v1/service"
	"starter-go-gin/utils"
)

// AuthService is a service for auth
type AuthService struct {
	cfg                config.Config
	authRepo           repository.AuthRepositoryUseCase
	userFinderService  service.UserFinderUseCase
	userUpdaterService service.UserUpdaterUseCase
	cache              interfaces.Cacheable
	cloudStorage       interfaces.CloudStorageUseCase
}

// AuthUseCase is a usecase for auth
type AuthUseCase interface {
	// AuthValidate is a function that validates the user
	AuthValidate(ctx context.Context, username, password, roleName string) (*entity.User, error)
	// AuthValidateAdmin is a function that validates the user
	AuthValidateAdmin(ctx context.Context, username, password string) (*entity.User, error)
	// GenerateAccessToken is a function that generates an access token
	GenerateAccessToken(ctx context.Context, user *entity.User) (*entity.Token, error)
	// GenerateAccessTokenAdmin is a function that generates an access token
	GenerateAccessTokenAdmin(ctx context.Context, user *entity.User) (*entity.Token, error)
	// Logout is a function that logs out the user
	Logout(ctx context.Context, userID uuid.UUID, deviceID string, merchantID uuid.UUID) error
}

// NewAuthService is a constructor for AuthService
func NewAuthService(
	cfg *config.Config,
	authRepo repository.AuthRepositoryUseCase,
	userFinderService service.UserFinderUseCase,
	cache interfaces.Cacheable,
	cloudStorage interfaces.CloudStorageUseCase,
) *AuthService {
	return &AuthService{
		cfg:               *cfg,
		authRepo:          authRepo,
		userFinderService: userFinderService,
		cache:             cache,
		cloudStorage:      cloudStorage,
	}
}

// AuthValidate is a function that validates the user
func (as *AuthService) AuthValidate(ctx context.Context, username, password, roleName string) (*entity.User, error) {
	user, err := as.userFinderService.GetUserByUsername(ctx, username, roleName)

	if err != nil && err.Error() != errors.ErrRecordNotFound.Error().Error() {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, errors.ErrLoginFailed.Error()
	}

	return user, nil
}

// AuthValidateAdmin is a function that validates the user
func (as *AuthService) AuthValidateAdmin(ctx context.Context, email, password string) (*entity.User, error) {
	admin, err := as.authRepo.GetUserByUsernameAndRole(ctx, email, constant.RoleAdmin)

	if err != nil {
		return nil, err
	}

	if admin == nil {
		return nil, errors.ErrLoginNotFound.Error()
	}

	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))

	if err != nil {
		return nil, errors.ErrLoginFailed.Error()
	}

	return admin, nil
}

// GenerateAccessToken is a function that generates an access token
func (as *AuthService) GenerateAccessToken(ctx context.Context, user *entity.User) (*entity.Token, error) {
	result := &entity.Token{}

	dataBytes, _ := as.cache.Get(fmt.Sprintf(commonCache.TokenUserByJTI, user.ID))

	if dataBytes != nil {
		if err := json.Unmarshal(dataBytes, &result); err != nil {
			return nil, err
		}

		decoded, err := utils.JWTDecode(as.cfg, result.Token)
		if err != nil {
			return nil, err
		}

		// return from cache if token is not expired
		if decoded.ExpiresAt > time.Now().Unix() {
			return result, nil
		}
	}

	userRole, err := as.userFinderService.GetUserByID(ctx, user.ID)

	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	bodyJWT := map[string]interface{}{
		"id":      user.ID,
		"role_id": userRole.RoleID,
		"name":    user.Username,
	}

	timeExp := time.Now().Add(time.Hour * constant.TwentyFourHour).Unix()

	token, _, err := utils.JWTEncode(as.cfg, bodyJWT, as.cfg.JWTConfig.Issuer, timeExp)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	result.Token = token

	if err := as.cache.Set(fmt.Sprintf(commonCache.TokenUserByJTI, user.ID), &result, commonCache.OneYear); err != nil {
		logger.Error(ctx, err)
	}

	return result, nil
}

// GenerateAccessTokenAdmin is a function that generates an access token
func (as *AuthService) GenerateAccessTokenAdmin(ctx context.Context, user *entity.User) (*entity.Token, error) {
	result := &entity.Token{}

	dataBytes, _ := as.cache.Get(fmt.Sprintf(commonCache.TokenUserByJTI, user.ID))

	if dataBytes != nil {
		if err := json.Unmarshal(dataBytes, &result); err != nil {
			return nil, err
		}

		decoded, err := utils.JWTDecode(as.cfg, result.Token)
		if err != nil {
			logger.ErrorWithStr(ctx, "error decoding token", err)
			return nil, err
		}

		// return from cache if token is not expired
		if decoded.ExpiresAt > time.Now().Unix() {
			return result, nil
		}
	}

	userRole, err := as.userFinderService.GetUserByID(ctx, user.ID)

	if err != nil {
		logger.Error(ctx, err)
		return nil, err
	}

	bodyJWT := map[string]interface{}{
		"id":      user.ID,
		"role_id": userRole.RoleID,
		"name":    user.Username,
	}

	timeExp := time.Now().Add(time.Hour * constant.TwentyFourHour).Unix()

	token, _, err := utils.JWTEncode(as.cfg, bodyJWT, as.cfg.JWTConfig.IssuerAdmin, timeExp)

	if err != nil {
		logger.ErrorWithStr(ctx, "error generating token", err)
		return nil, errors.ErrInternalServerError.Error()
	}

	result.Token = token

	if err := as.cache.Set(fmt.Sprintf(commonCache.TokenUserByJTI, user.ID), &result, commonCache.OneYear); err != nil {
		logger.Error(ctx, err)
	}

	return result, nil
}

// Logout is a function that logs out the user
func (as *AuthService) Logout(ctx context.Context, userID uuid.UUID, deviceID string, merchantID uuid.UUID) error {
	// Remove token from redis
	if err := as.cache.BulkRemove(fmt.Sprintf(commonCache.TokenUserByJTI, userID)); err != nil {
		logger.ErrorWithStr(ctx, "[AuthService - Logout] ", err)
	}
	return nil
}
