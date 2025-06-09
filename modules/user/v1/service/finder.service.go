package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/user/v1/repository"
)

// UserFinder is a service for user
type UserFinder struct {
	ufg          config.Config
	userRepo     repository.UserRepositoryUseCase
	userRoleRepo repository.UserRoleRepositoryUseCase
	roleRepo     repository.RoleRepositoryUseCase
}

// UserFinderUseCase is a usecase for user
type UserFinderUseCase interface {
	// GetUsers gets all users
	GetUsers(ctx context.Context, query, order, sort string, limit, offset int) ([]*entity.User, int64, error)
	// GetUserByID gets a user by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	// GetUserByEmail gets user by email
	GetUserByEmail(ctx context.Context, email, roleName string) (*entity.User, error)
	// GetUserByUsername gets user by username
	GetUserByUsername(ctx context.Context, username, roleName string) (*entity.User, error)
	// GetRoles gets all roles
	GetRoles(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Role, error)
	// FindUserRoleByUserID finds user role by user id
	FindUserRoleByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserRole, error)
	// GetRoleByID gets role by id
	GetRoleByID(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	// GetRoleByName is a function that gets role by name
	GetRoleByName(ctx context.Context, name string) (*entity.Role, error)
}

// NewUserFinder creates a new UserFinder
func NewUserFinder(
	ufg config.Config,
	userRepo repository.UserRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
) *UserFinder {
	return &UserFinder{
		ufg:          ufg,
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		roleRepo:     roleRepo,
	}
}

// GetAdminUserByID gets a admin user by ID
func (uf *UserFinder) GetAdminUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := uf.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	if user == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return user, nil
}

// GetUserByEmail gets user by email
func (uf *UserFinder) GetUserByEmail(ctx context.Context, email, roleName string) (*entity.User, error) {
	user, err := uf.userRepo.GetUserByEmail(ctx, email, roleName)

	if err != nil {
		return user, errors.ErrInternalServerError.Error()
	}

	if user == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return user, nil
}

// GetUserExistsByEmail gets user exists by email
func (uf *UserFinder) GetUserExistsByEmail(ctx context.Context, email string) ([]*entity.User, error) {
	user, err := uf.userRepo.GetUserExistsByEmail(ctx, email)

	if err != nil {
		return user, errors.ErrInternalServerError.Error()
	}

	if user == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return user, nil
}

// GetRoles gets all roles
func (uf *UserFinder) GetRoles(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Role, error) {
	roles, err := uf.roleRepo.FindAll(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return roles, nil
}

// GetRoleCMS gets role cms
func (uf *UserFinder) GetRoleCMS(ctx context.Context, query string) ([]*entity.Role, error) {
	roles, err := uf.roleRepo.FindByType(ctx, "cms", query)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return roles, nil
}

// GetUserByUsername gets user by phone number
func (uf *UserFinder) GetUserByUsername(ctx context.Context, username, roleName string) (*entity.User, error) {
	user, err := uf.userRepo.GetUserByUsername(ctx, username, roleName)

	if err != nil {
		return user, errors.ErrInternalServerError.Error()
	}

	if user == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}
	return user, nil
}
