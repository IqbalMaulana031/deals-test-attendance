package service

import (
	"context"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/master/v1/repository"

	"github.com/google/uuid"
)

// MasterFinder is a service for product
type MasterFinder struct {
	cfg             config.Config
	shiftRepo       repository.ShiftRepositoryUseCase
	shiftDetailRepo repository.ShiftDetailRepositoryUseCase
}

// MasterFinderUseCase is a usecase for product
type MasterFinderUseCase interface {
	// GetShiftByID gets a shift by ID
	GetShiftByID(ctx context.Context, ID uuid.UUID) (*entity.Shift, error)
	// GetShift finds all shifts by filter
	GetShift(ctx context.Context, query, sort, order string, limit, page int) ([]*entity.Shift, int64, error)
	// GetShiftAndDetailsByID finds a shift and its details by ID
	GetShiftAndDetailsByID(ctx context.Context, ID uuid.UUID) (*entity.Shift, error)
	// GetShiftDetailByID gets a shift detail by ID
	GetShiftDetailByID(ctx context.Context, ID uuid.UUID) (*entity.ShiftDetail, error)
}

// NewMasterFinder creates a new MasterFinder
func NewMasterFinder(
	cfg config.Config,
	ShiftRepo repository.ShiftRepositoryUseCase,
	shiftDetailRepo repository.ShiftDetailRepositoryUseCase,
) *MasterFinder {
	return &MasterFinder{
		cfg:             cfg,
		shiftRepo:       ShiftRepo,
		shiftDetailRepo: shiftDetailRepo,
	}
}
