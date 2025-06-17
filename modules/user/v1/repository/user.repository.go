package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	commonCache "starter-go-gin/common/cache"
	"starter-go-gin/common/constant"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/logger"
	"starter-go-gin/common/tools"
	"starter-go-gin/entity"
)

// UserRepository is a repository for user
type UserRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// UserRepositoryUseCase is a use case for user
type UserRepositoryUseCase interface {
	// GetUserByID is a function to get user by id
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	// Update is a function to update user
	Update(ctx context.Context, user *entity.User) error
	// CreateUser is a function to create user
	CreateUser(ctx context.Context, user *entity.User) error
	// GetUsers is a function to get users
	GetUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error)
	// GetUserByUsername is a function to get user by phone number
	GetUserByUsername(ctx context.Context, username, roleName string) (*entity.User, error)
	// DeleteUser is a function to delete user
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(db *gorm.DB, cache interfaces.Cacheable) *UserRepository {
	return &UserRepository{db: db, cache: cache}
}

// GetUserByID is a function to get user by id
func (ur *UserRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	result := &entity.User{}

	bytes, _ := ur.cache.Get(fmt.Sprintf(
		commonCache.UserByID, id.String()))

	if bytes != nil {
		if err := json.Unmarshal(bytes, &result); err != nil {
			return nil, err
		}
		return result, nil
	}

	if err := ur.db.
		WithContext(ctx).
		Preload("Role").
		Where("id = ?", id).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByID] user not found")
	}

	if err := ur.cache.Set(fmt.Sprintf(commonCache.UserByID, id), &result, commonCache.OneMonth); err != nil {
		return nil, err
	}

	return result, nil
}

// Update is a function to update user
func (ur *UserRepository) Update(ctx context.Context, user *entity.User) error {
	oldTime := user.UpdatedAt
	user.UpdatedAt = time.Now()
	if err := ur.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.User)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, user.ID).Error; err != nil {
				logger.Error(ctx, err)
				return err
			}
			if err := tx.Model(&entity.User{}).
				Where(`id`, user.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(user)).Error; err != nil {
				logger.Error(ctx, err)
				return err
			}
			return nil
		}); err != nil {
		user.UpdatedAt = oldTime
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.UserByID, user.ID)); err != nil {
		return err
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.UserFindByUSername, user.Username)); err != nil {
		return err
	}

	return nil
}

// CreateUser is a function to create user
func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	if err := ur.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Create(user).
		Error; err != nil {
		return errors.Wrap(err, "[UserRepository-CreateUser] error while creating user")
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.UserFindByUSername, user.Username)); err != nil {
		return err
	}

	return nil
}

// GetUsers is a function to get all users
func (ur *UserRepository) GetUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error) {
	var user []*entity.User
	var total int64
	var gormDB = ur.db.
		WithContext(ctx).
		Model(&entity.User{}).
		Joins("join auth.roles r on users.role_id = r.id").
		Where(`r."name" = ?`, constant.EmployeeRoleName)

	gormDB.Count(&total)

	gormDB = gormDB.Limit(limit).
		Offset(offset * limit)

	if query != "" {
		gormDB = gormDB.
			Where("name ILIKE ?", "%"+query+"%").
			Or("email ILIKE ?", "%"+query+"%").
			Or("phone_number ILIKE ?", "%"+query+"%")
	}

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", tools.EscapeSpecial(sort), tools.EscapeSpecial(order)))

	if err := gormDB.Find(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[UserRepository-GetUsers] error when looking up all user")
	}

	return user, total, nil
}

// GetUserByUsername is a function to get user by phone number
func (ur *UserRepository) GetUserByUsername(ctx context.Context, username, roleName string) (*entity.User, error) {
	user := &entity.User{}

	bytes, _ := ur.cache.Get(fmt.Sprintf(commonCache.UserFindByUSername, username))
	if bytes != nil {
		if err := json.Unmarshal(bytes, &user); err != nil {
			return nil, err
		}
		return user, nil
	}

	if err := ur.db.WithContext(ctx).
		Joins("join auth.roles r on auth.users.role_id = r.id").
		Where(`r."name" = ?`, constant.EmployeeRoleName).
		Where(`username = ?`, username).
		First(&user).Error; err != nil {
		fmt.Println(err)
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[UserRepository-GetUserByUsername] error when looking up user")
	}

	if err := ur.cache.Set(fmt.Sprintf(commonCache.UserFindByUSername, username), &user, commonCache.OneMonth); err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser is a function to delete user
func (ur *UserRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := ur.db.WithContext(ctx).
		Model(&entity.User{}).
		Where(`id = ?`, id).
		Delete(&entity.User{}, "id = ?", id).Error; err != nil {
		return errors.Wrap(err, "[UserRepository-DeleteUser] error when updating user data")
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.UserByID, "*")); err != nil {
		return err
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.UserFindByUSername, "*")); err != nil {
		return err
	}

	return nil
}
