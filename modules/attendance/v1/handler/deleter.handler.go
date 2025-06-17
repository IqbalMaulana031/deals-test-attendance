package handler

import (
	"starter-go-gin/common/interfaces"
	"starter-go-gin/config"
	"starter-go-gin/modules/attendance/v1/service"
)

// AttendanceDeleterHandler is a handler for attendance finder
type AttendanceDeleterHandler struct {
	cfg               config.Config
	attendanceDeleter service.AttendanceDeleterUseCase
	cache             interfaces.Cacheable
}

// NewAttendanceDeleterHandler is a constructor for AttendanceDeleterHandler
func NewAttendanceDeleterHandler(
	cfg config.Config,
	attendanceDeleter service.AttendanceDeleterUseCase,
) *AttendanceDeleterHandler {
	return &AttendanceDeleterHandler{
		cfg:               cfg,
		attendanceDeleter: attendanceDeleter,
	}
}
