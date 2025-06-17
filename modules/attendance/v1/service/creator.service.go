package service

import (
	"context"
	"starter-go-gin/config"
	"starter-go-gin/modules/attendance/v1/repository"
	masterRepo "starter-go-gin/modules/master/v1/repository"

	"github.com/google/uuid"
)

// AttendanceCreator is a struct that contains all the dependencies for the Attendance creator
type AttendanceCreator struct {
	cfg                  config.Config
	AttendanceRepository repository.AttendanceRepositoryUseCase
	ShiftDetailRepo      masterRepo.ShiftDetailRepositoryUseCase
}

// AttendanceCreatorUseCase is a use case for the Attendance creator
type AttendanceCreatorUseCase interface {
	// RechordTime employee's attendance
	RechordTime(ctx context.Context, userId, shiftDetailID uuid.UUID) error
}

// NewAttendanceCreator is a constructor for the Attendance creator
func NewAttendanceCreator(
	cfg config.Config,
	AttendanceRepository repository.AttendanceRepositoryUseCase,
	ShiftDetailRepo masterRepo.ShiftDetailRepositoryUseCase,
) *AttendanceCreator {
	return &AttendanceCreator{
		cfg:                  cfg,
		AttendanceRepository: AttendanceRepository,
		ShiftDetailRepo:      ShiftDetailRepo,
	}
}
