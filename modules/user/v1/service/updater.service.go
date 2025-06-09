package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/user/v1/repository"
)

// UserUpdater is a struct that contains the dependencies of UserUpdater
type UserUpdater struct {
	cfg          config.Config
	userRepo     repository.UserRepositoryUseCase
	userRoleRepo repository.UserRoleRepositoryUseCase
	roleRepo     repository.RoleRepositoryUseCase
}

// UserUpdaterUseCase is a struct that contains the dependencies of UserUpdaterUseCase
type UserUpdaterUseCase interface {
	// Update a user
	// UpdateRole updates a role
	UpdateRole(ctx context.Context, id uuid.UUID, name string, permissionIDs []uuid.UUID) error
	// UpdateUser updates an employee.
	UpdateUser(ctx context.Context, user *entity.User) error
	// UpdateUserRole updates a user role
	UpdateUserRole(ctx context.Context, userID, roleID uuid.UUID) error
}

// NewUserUpdater is a function that creates a new UserUpdater
func NewUserUpdater(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
) *UserUpdater {
	return &UserUpdater{
		cfg:          cfg,
		userRepo:     userRepo,
		userRoleRepo: userRoleRepo,
		roleRepo:     roleRepo,
	}
}

// UpdateRole updates a role
func (uu *UserUpdater) UpdateRole(ctx context.Context, id uuid.UUID, name string, permissionIDs []uuid.UUID) error {
	role, err := uu.roleRepo.FindByID(ctx, id)

	if err != nil {
		return errors.ErrInternalServerError.Error()
	}

	if role == nil {
		return errors.ErrRecordNotFound.Error()
	}

	role.Label = name

	if err := uu.roleRepo.Update(ctx, role); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
