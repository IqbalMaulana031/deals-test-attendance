package service

import (
	"context"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/master/v1/repository"
)

// MasterCreator is a struct that contains all the dependencies for the Master creator
type MasterCreator struct {
	cfg             config.Config
	shiftRepo       repository.ShiftRepositoryUseCase
	shiftDetailRepo repository.ShiftDetailRepositoryUseCase
}

// MasterCreatorUseCase is a use case for the Master creator
type MasterCreatorUseCase interface {
	// CreateShift creates a new shift
	CreateShift(ctx context.Context, shift *entity.Shift, shiftDetail *[]entity.ShiftDetail) error
}

// NewMasterCreator is a constructor for the Master creator
func NewMasterCreator(
	cfg config.Config,
	shiftRepo repository.ShiftRepositoryUseCase,
	shiftDetailRepo repository.ShiftDetailRepositoryUseCase,
) *MasterCreator {
	return &MasterCreator{
		cfg:             cfg,
		shiftRepo:       shiftRepo,
		shiftDetailRepo: shiftDetailRepo,
	}
}
