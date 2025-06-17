package handler

import (
	"starter-go-gin/config"
	"starter-go-gin/modules/attendance/v1/service"
)

// AttendanceUpdaterHandler is a handler for product updater
type AttendanceUpdaterHandler struct {
	attendanceUpdater service.AttendanceUpdaterUseCase
	cfg               config.Config
}

// NewAttendanceUpdaterHandler is a constructor for AttendanceUpdaterHandler
func NewAttendanceUpdaterHandler(
	attendanceUpdater service.AttendanceUpdaterUseCase,
	cfg config.Config,
) *AttendanceUpdaterHandler {
	return &AttendanceUpdaterHandler{
		attendanceUpdater: attendanceUpdater,
		cfg:               cfg,
	}
}
