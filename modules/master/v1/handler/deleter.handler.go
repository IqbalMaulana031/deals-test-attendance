package handler

import (
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/modules/master/v1/service"
)

// MasterDeleterHandler is a handler for master finder
type MasterDeleterHandler struct {
	cfg           config.Config
	masterDeleter service.MasterDeleterUseCase
	masterFinder  service.MasterFinderUseCase
	cache         interfaces.Cacheable
}

// NewMasterDeleterHandler is a constructor for MasterDeleterHandler
func NewMasterDeleterHandler(
	cfg config.Config,
	masterDeleter service.MasterDeleterUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
	masterFinder service.MasterFinderUseCase,
	cacheable interfaces.Cacheable,
) *MasterDeleterHandler {
	return &MasterDeleterHandler{
		cfg:           cfg,
		masterDeleter: masterDeleter,
		masterFinder:  masterFinder,
		cache:         cacheable,
	}
}
