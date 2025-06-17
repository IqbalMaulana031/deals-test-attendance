package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/user/v1/repository"
)

// UserFinder is a service for user
type UserFinder struct {
	ufg      config.Config
	userRepo repository.UserRepositoryUseCase
	roleRepo repository.RoleRepositoryUseCase
}

// UserFinderUseCase is a usecase for user
type UserFinderUseCase interface {
	// GetUsers gets all users
	GetUsers(ctx context.Context, query, order, sort string, limit, offset int) ([]*entity.User, int64, error)
	// GetUserByID gets a user by ID
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	// GetUserByUsername gets user by username
	GetUserByUsername(ctx context.Context, username, roleName string) (*entity.User, error)
	// GetRoles gets all roles
	GetRoles(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Role, error)
	// GetRoleByID gets role by id
	GetRoleByID(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	// GetRoleByName is a function that gets role by name
	GetRoleByName(ctx context.Context, name string) (*entity.Role, error)
}

// NewUserFinder creates a new UserFinder
func NewUserFinder(
	ufg config.Config,
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
) *UserFinder {
	return &UserFinder{
		ufg:      ufg,
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}
