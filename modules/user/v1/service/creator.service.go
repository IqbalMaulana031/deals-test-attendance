package service

import (
	"context"

	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/user/v1/repository"
)

// UserCreator is a struct that contains all the dependencies for the User creator
type UserCreator struct {
	cfg          config.Config
	userRepo     repository.UserRepositoryUseCase
	roleRepo     repository.RoleRepositoryUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// UserCreatorUseCase is a use case for the User creator
type UserCreatorUseCase interface {
	// CreateUser creates a new user
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

// NewUserCreator is a constructor for the User creator
func NewUserCreator(
	cfg config.Config,
	userRepo repository.UserRepositoryUseCase,
	roleRepo repository.RoleRepositoryUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserCreator {
	return &UserCreator{
		cfg:          cfg,
		userRepo:     userRepo,
		roleRepo:     roleRepo,
		cloudStorage: cloudStorage,
	}
}
