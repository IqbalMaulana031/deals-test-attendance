package service

import (
	"context"
	"errors"
	"log"
	"starter-go-gin/entity"
	"time"
)

func (au *AttendanceUpdater) UpdateAttendance(ctx context.Context, attendance *entity.Attendance) error {
	err := au.attendanceRepo.UpdateAttendance(ctx, &entity.Attendance{
		ShiftID:  attendance.ID,
		Checkin:  attendance.Checkin,
		Checkout: attendance.Checkout,
		Auditable: entity.Auditable{
			UpdatedAt: time.Now(),
			UpdatedBy: attendance.UpdatedBy,
		},
	}, attendance.ID)
	if err != nil {
		log.Println("[MasterUpdater-UpdateAttendance] Error updating attendance:", err)
		return errors.New("failed to update attendance")
	}
	return nil
}
