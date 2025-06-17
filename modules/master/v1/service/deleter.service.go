package service

import (
	"context"
	"starter-go-gin/config"
	"starter-go-gin/modules/master/v1/repository"

	"github.com/google/uuid"
)

// MasterDeleter is a service for master
type MasterDeleter struct {
	cfg             config.Config
	shiftRepo       repository.ShiftRepositoryUseCase
	shiftDetailRepo repository.ShiftDetailRepositoryUseCase
}

// MasterDeleterUseCase is a use case for master
type MasterDeleterUseCase interface {
	// delete shift
	DeleteShift(ctx context.Context, id uuid.UUID) error
}

// NewMasterDeleter creates a new MasterDeleter
func NewMasterDeleter(
	cfg config.Config,
	shiftRepo repository.ShiftRepositoryUseCase,
	shiftDetailRepo repository.ShiftDetailRepositoryUseCase,
) *MasterDeleter {
	return &MasterDeleter{
		cfg:             cfg,
		shiftRepo:       shiftRepo,
		shiftDetailRepo: shiftDetailRepo,
	}
}
