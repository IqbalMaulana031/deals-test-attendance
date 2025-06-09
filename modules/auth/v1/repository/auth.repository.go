package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	commonCache "starter-go-gin/common/cache"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/logger"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/utils"
)

// AuthRepository is a repository for auth
type AuthRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
	// The code below is commented out because synchronization with Firestore is no longer needed
	// firestore interfaces.FirestoreUseCase
	cfg config.Config
}

// AuthRepositoryUseCase is a repository for auth
type AuthRepositoryUseCase interface {
	// GetUserByID finds an user by id
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	// GetUserByEmailAndUserType finds a user by email and user type
	GetUserByEmailAndUserType(ctx context.Context, email, userType string) (*entity.User, error)
	// 	UpdateForgotPasswordToken updates forgot password token
	UpdateForgotPasswordToken(ctx context.Context, user *entity.User, token string) error
	// UpdateUser updates user
	UpdateUser(ctx context.Context, user *entity.User) error
	// Register creates a new user
	Register(ctx context.Context, user *entity.User, userRole *entity.UserRole) error
	// GetUserByForgotPasswordToken finds a user by forgot password token
	GetUserByForgotPasswordToken(ctx context.Context, token string) (*entity.User, error)
}

// NewAuthRepository returns a auth repository
func NewAuthRepository(
	db *gorm.DB,
	cache interfaces.Cacheable,
	// The code below is commented out because synchronization with Firestore is no longer needed
	// firestore interfaces.FirestoreUseCase,
	cfg config.Config,
) *AuthRepository {
	return &AuthRepository{db, cache, cfg}
}

// GetUserByID finds a user by id
func (ar *AuthRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	result := &entity.User{}

	if err := ar.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logger.Error(ctx, errors.Wrap(err, "[AuthRepository-GetUserByID] id not found"))
		return nil, errors.Wrap(err, "[UserRepository-GetUserByID] id not found")
	}

	return result, nil
}

// UpdateForgotPasswordToken updates forgot password token
func (ar *AuthRepository) UpdateForgotPasswordToken(ctx context.Context, user *entity.User, token string) error {
	oldTime := user.UpdatedAt
	user.UpdatedAt = time.Now()
	if err := ar.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.User)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, user.ID).Error; err != nil {
				return errors.Wrap(err, "[UserRepository-ChangePassword] error when updating data")
			}
			if err := tx.Model(&entity.User{}).
				Where(`id = ?`, user.ID).Update("forgot_password_token", utils.StringToNullString(token)).Error; err != nil {
				return errors.Wrap(err, "[UserRepository-Update] error when updating data User b")
			}
			return nil
		}); err != nil {
		user.UpdatedAt = oldTime
	}
	return nil
}

// GetUserByEmailAndUserType finds a user by email and user type
func (ar *AuthRepository) GetUserByEmailAndUserType(ctx context.Context, email, userType string) (*entity.User, error) {
	result := new(entity.User)

	if err := ar.db.
		WithContext(ctx).
		Joins("inner join auth.user_roles on auth.users.id = auth.user_roles.user_id").
		Joins("inner join auth.roles  on auth.user_roles.role_id = auth.roles.id").
		Where("email = ?", email).
		Where("auth.roles.name = ?", userType).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetAdminByEmail] email not found")
	}

	return result, nil
}

// Register creates a new user
func (ar *AuthRepository) Register(ctx context.Context, user *entity.User, userRole *entity.UserRole) error {
	if err := ar.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// create user
		if err := tx.Model(&entity.User{}).Create(user).Error; err != nil {
			return errors.Wrap(err, "[UserRepository-Register] error when creating data User")
		}

		// create user role
		if err := tx.Model(&entity.UserRole{}).Create(userRole).Error; err != nil {
			return errors.Wrap(err, "[UserRepository-Register] error when creating data UserRole")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "[UserRepository-Register] error when creating data")
	}
	return nil
}

// GetUserByForgotPasswordToken finds a user by forgot password token
func (ar *AuthRepository) GetUserByForgotPasswordToken(ctx context.Context, token string) (*entity.User, error) {
	result := &entity.User{}

	if err := ar.db.
		WithContext(ctx).
		Where("forgot_password_token = ?", token).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByForgotPasswordToken] token not found")
	}

	return result, nil
}

// UpdateUser updates user
func (ar *AuthRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	oldTime := user.UpdatedAt
	user.UpdatedAt = time.Now()
	if err := ar.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.User)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, user.ID).Error; err != nil {
				return errors.Wrap(err, "[UserRepository-UpdateUser] error when updating data")
			}
			if err := tx.Model(&entity.User{}).
				Where(`id = ?`, user.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(user)).Error; err != nil {
				return errors.Wrap(err, "[UserRepository-Update] error when updating data User b")
			}
			return nil
		}); err != nil {
		user.UpdatedAt = oldTime
	}

	if err := ar.cache.BulkRemove(fmt.Sprintf(commonCache.UserByID, user.ID)); err != nil {
		return err
	}

	if err := ar.cache.BulkRemove(fmt.Sprintf(commonCache.UserFindByEmail, user.Email)); err != nil {
		return err
	}

	// if err := ar.cache.BulkRemove(fmt.Sprintf(commonCache.UserFindByPhoneNumber, user.PhoneNumber)); err != nil {
	// 	return err
	// }

	return nil
}
