package handler

import (
	"starter-go-gin/config"
	"starter-go-gin/modules/utils/v1/service"
)

// UtilsCreatorHandler is a handler for utils
type UtilsCreatorHandler struct {
	cfg     config.Config
	creator service.UtilsCreatorUseCase
}

// NewUtilsCreatorHandler is a constructor for UtilsCreatorHandler
func NewUtilsCreatorHandler(
	cfg config.Config,
	creator service.UtilsCreatorUseCase,
) *UtilsCreatorHandler {
	return &UtilsCreatorHandler{
		cfg:     cfg,
		creator: creator,
	}
}
