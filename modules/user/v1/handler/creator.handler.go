package handler

import (
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/modules/user/v1/service"
)

// UserCreatorHandler is a handler for user finder
type UserCreatorHandler struct {
	cfg          config.Config
	userCreator  service.UserCreatorUseCase
	cloudStorage interfaces.CloudStorageUseCase
	userFinder   service.UserFinderUseCase
	cache        interfaces.Cacheable
}

// NewUserCreatorHandler is a constructor for UserCreatorHandler
func NewUserCreatorHandler(
	cfg config.Config,
	userCreator service.UserCreatorUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
	userFinder service.UserFinderUseCase,
	cache interfaces.Cacheable,
) *UserCreatorHandler {
	return &UserCreatorHandler{
		cfg:          cfg,
		userCreator:  userCreator,
		cloudStorage: cloudStorage,
		userFinder:   userFinder,
		cache:        cache,
	}
}
