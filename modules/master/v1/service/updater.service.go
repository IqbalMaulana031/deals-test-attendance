package service

import (
	"context"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/master/v1/repository"
)

// MasterUpdater is a struct that contains the dependencies of MasterUpdater
type MasterUpdater struct {
	cfg             config.Config
	shiftRepo       repository.ShiftRepositoryUseCase
	shiftDetailRepo repository.ShiftDetailRepositoryUseCase
}

// MasterUpdaterUseCase is a struct that contains the dependencies of MasterUpdaterUseCase
type MasterUpdaterUseCase interface {
	// UpdateShift updates an existing shift
	UpdateShift(ctx context.Context, shift *entity.Shift) error
	// UpdateShiftDetail updates an existing shift detail
	UpdateShiftDetail(ctx context.Context, shiftDetail *entity.ShiftDetail) error
}

// NewMasterUpdater is a function that creates a new MasterUpdater
func NewMasterUpdater(
	cfg config.Config,
	shiftRepo repository.ShiftRepositoryUseCase,
	shiftDetailRepo repository.ShiftDetailRepositoryUseCase,
) *MasterUpdater {
	return &MasterUpdater{
		cfg:             cfg,
		shiftRepo:       shiftRepo,
		shiftDetailRepo: shiftDetailRepo,
	}
}
