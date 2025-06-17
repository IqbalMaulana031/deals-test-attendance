package service

import (
	"context"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/modules/attendance/v1/repository"
)

// AttendanceUpdater is a struct that contains the dependencies of AttendanceUpdater
type AttendanceUpdater struct {
	cfg            config.Config
	attendanceRepo repository.AttendanceRepositoryUseCase
}

// AttendanceUpdaterUseCase is a struct that contains the dependencies of AttendanceUpdaterUseCase
type AttendanceUpdaterUseCase interface {
	// UpdateAttendance updates an existing attendance
	UpdateAttendance(ctx context.Context, attendance *entity.Attendance) error
}

// NewAttendanceUpdater is a function that creates a new AttendanceUpdater
func NewAttendanceUpdater(
	cfg config.Config,
	attendanceRepo repository.AttendanceRepositoryUseCase,
) *AttendanceUpdater {
	return &AttendanceUpdater{
		cfg:            cfg,
		attendanceRepo: attendanceRepo,
	}
}
