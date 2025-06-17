package service

import (
	"context"

	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/user/v1/repository"
)

// UserUpdater is a struct that contains the dependencies of UserUpdater
type UserUpdater struct {
	cfg      config.Config
	userRepo repository.UserRepositoryUseCase
	roleRepo repository.RoleRepositoryUseCase
}

// UserUpdaterUseCase is a struct that contains the dependencies of UserUpdaterUseCase
type UserUpdaterUseCase interface {
	// Update a user
	// UpdateUser updates an employee.
	UpdateUser(ctx context.Context, user *entity.User) error
}

// NewUserUpdater is a function that creates a new UserUpdater
func NewUserUpdater(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
) *UserUpdater {
	return &UserUpdater{
		cfg:      cfg,
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}
