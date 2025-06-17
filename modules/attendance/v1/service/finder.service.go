package service

import (
	"context"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/attendance/v1/repository"

	"github.com/google/uuid"
)

// AttendanceFinder is a service for attendance
type AttendanceFinder struct {
	cfg            config.Config
	attendanceRepo repository.AttendanceRepositoryUseCase
}

// AttendanceFinderUseCase is a usecase for attendance
type AttendanceFinderUseCase interface {
	// GetAttendanceByID gets a attendance by ID
	GetAttendanceByID(ctx context.Context, ID uuid.UUID) (*entity.Attendance, error)
	// GetAttendance finds all attendances by filter
	GetAttendance(ctx context.Context, query, sort, order string, limit, page int) ([]*entity.Attendance, int64, error)
	// GetAttendanceByUserID gets a attendance by user ID
	GetAttendanceByUserID(ctx context.Context, userID uuid.UUID, query, sort, order string, limit, page int) ([]*entity.Attendance, int64, error)
}

// NewAttendanceFinder creates a new AttendanceFinder
func NewAttendanceFinder(
	cfg config.Config,
	AttendanceRepo repository.AttendanceRepositoryUseCase,
) *AttendanceFinder {
	return &AttendanceFinder{
		cfg:            cfg,
		attendanceRepo: AttendanceRepo,
	}
}
