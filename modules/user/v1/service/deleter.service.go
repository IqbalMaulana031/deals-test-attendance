package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/config"
	"starter-go-gin/modules/user/v1/repository"
)

// UserDeleter is a service for user
type UserDeleter struct {
	cfg          config.Config
	userRepo     repository.UserRepositoryUseCase
	roleRepo     repository.RoleRepositoryUseCase
	userRoleRepo repository.UserRoleRepositoryUseCase
}

// UserDeleterUseCase is a use case for user
type UserDeleterUseCase interface {
	// DeleteRole deletes role
	DeleteRole(ctx context.Context, id uuid.UUID, deletedBy string) error
	// DeleteUser deletes user
	DeleteUser(ctx context.Context, id, merchantID uuid.UUID) error
}

// NewUserDeleter creates a new UserDeleter
func NewUserDeleter(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
	userRoleRepo repository.UserRoleRepositoryUseCase,
) *UserDeleter {
	return &UserDeleter{
		cfg:          cfg,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		userRoleRepo: userRoleRepo,
	}
}

// DeleteRole deletes role
func (ud *UserDeleter) DeleteRole(ctx context.Context, id uuid.UUID, deletedBy string) error {
	roleIDs := []uuid.UUID{id}

	checkRole, err := ud.userRoleRepo.FindUserRoleByRoleIDs(ctx, roleIDs)

	if err != nil {
		return errors.ErrInternalServerError.Error()
	}

	if len(checkRole) > 0 {
		return errors.ErrCannotDeleteRole.Error()
	}

	if err := ud.roleRepo.Delete(ctx, id, deletedBy); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
