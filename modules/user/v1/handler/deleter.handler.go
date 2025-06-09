package handler

import (
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/modules/user/v1/service"
)

// UserDeleterHandler is a handler for user finder
type UserDeleterHandler struct {
	cfg          config.Config
	userDeleter  service.UserDeleterUseCase
	cloudStorage interfaces.CloudStorageUseCase
	userFinder   service.UserFinderUseCase
	cache        interfaces.Cacheable
}

// NewUserDeleterHandler is a constructor for UserDeleterHandler
func NewUserDeleterHandler(
	cfg config.Config,
	userDeleter service.UserDeleterUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
	userFinder service.UserFinderUseCase,
	cacheable interfaces.Cacheable,
) *UserDeleterHandler {
	return &UserDeleterHandler{
		cfg:          cfg,
		userDeleter:  userDeleter,
		cloudStorage: cloudStorage,
		userFinder:   userFinder,
		cache:        cacheable,
	}
}
