package handler

import (
	"starter-go-gin/config"
	"starter-go-gin/modules/master/v1/service"
)

// MasterUpdaterHandler is a handler for product updater
type MasterUpdaterHandler struct {
	masterUpdater service.MasterUpdaterUseCase
	cfg           config.Config
}

// NewMasterUpdaterHandler is a constructor for MasterUpdaterHandler
func NewMasterUpdaterHandler(
	masterUpdater service.MasterUpdaterUseCase,
	cfg config.Config,
) *MasterUpdaterHandler {
	return &MasterUpdaterHandler{
		masterUpdater: masterUpdater,
		cfg:           cfg,
	}
}
