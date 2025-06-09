package handler

import (
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/modules/user/v1/service"
)

// UserFinderHandler is a handler for user finder
type UserFinderHandler struct {
	userFinder   service.UserFinderUseCase
	cfg          config.Config
	cache        interfaces.Cacheable
	gotenberg    interfaces.GotenbergUseCase
	excelize     interfaces.ExcelizeUseCase
	cloudStorage interfaces.CloudStorageUseCase
}

// NewUserFinderHandler is a constructor for UserFinderHandler
func NewUserFinderHandler(
	userFinder service.UserFinderUseCase,
	cache interfaces.Cacheable,
	cfg config.Config,
	gotenberg interfaces.GotenbergUseCase,
	excelize interfaces.ExcelizeUseCase,
	cloudStorage interfaces.CloudStorageUseCase,
) *UserFinderHandler {
	return &UserFinderHandler{
		userFinder:   userFinder,
		cache:        cache,
		gotenberg:    gotenberg,
		excelize:     excelize,
		cloudStorage: cloudStorage,
	}
}
