package service

import (
	"context"
	"starter-go-gin/config"
	"starter-go-gin/modules/attendance/v1/repository"

	"github.com/google/uuid"
)

// AttendanceDeleter is a service for attendance
type AttendanceDeleter struct {
	cfg            config.Config
	AttendanceRepo repository.AttendanceRepositoryUseCase
}

// AttendanceDeleterUseCase is a use case for attendance
type AttendanceDeleterUseCase interface {
	// delete Attendance
	DeleteAttendance(ctx context.Context, id uuid.UUID) error
}

// NewAttendanceDeleter creates a new AttendanceDeleter
func NewAttendanceDeleter(
	cfg config.Config,
	AttendanceRepo repository.AttendanceRepositoryUseCase,
) *AttendanceDeleter {
	return &AttendanceDeleter{
		cfg:            cfg,
		AttendanceRepo: AttendanceRepo,
	}
}
