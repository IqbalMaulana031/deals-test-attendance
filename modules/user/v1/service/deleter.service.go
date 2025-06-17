package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/config"
	"starter-go-gin/modules/user/v1/repository"
)

// UserDeleter is a service for user
type UserDeleter struct {
	cfg      config.Config
	userRepo repository.UserRepositoryUseCase
	roleRepo repository.RoleRepositoryUseCase
}

// UserDeleterUseCase is a use case for user
type UserDeleterUseCase interface {
	// DeleteUser deletes user
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

// NewUserDeleter creates a new UserDeleter
func NewUserDeleter(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
) *UserDeleter {
	return &UserDeleter{
		cfg:      cfg,
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}
