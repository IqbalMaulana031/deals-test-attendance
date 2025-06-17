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
	// GetUserByUsernameAndRole finds a user by email and user type
	GetUserByUsernameAndRole(ctx context.Context, email, userType string) (*entity.User, error)
	// UpdateUser updates user
	UpdateUser(ctx context.Context, user *entity.User) error
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

// GetUserByUsernameAndRole finds a user by username and user role
func (ar *AuthRepository) GetUserByUsernameAndRole(ctx context.Context, username, role string) (*entity.User, error) {
	result := new(entity.User)

	if err := ar.db.
		WithContext(ctx).
		Joins("join auth.roles r on users.role_id = r.id").
		Where(`r."name" = ?`, role).
		Where(`username = ?`, username).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetAdminByUsername] username not found")
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

	if err := ar.cache.BulkRemove(fmt.Sprintf(commonCache.UserFindByUSername, user.Username)); err != nil {
		return err
	}

	// if err := ar.cache.BulkRemove(fmt.Sprintf(commonCache.UserFindByPhoneNumber, user.PhoneNumber)); err != nil {
	// 	return err
	// }

	return nil
}
