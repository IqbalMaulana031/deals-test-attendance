package service

import (
	"context"
	"log"
	"starter-go-gin/common/errors"
	"starter-go-gin/entity"

	"github.com/google/uuid"
)

func (af *AttendanceFinder) GetAttendanceByID(ctx context.Context, ID uuid.UUID) (*entity.Attendance, error) {
	attendance, err := af.attendanceRepo.GetAttendanceByID(ctx, ID)
	if err != nil {
		log.Fatal("[AttendanceFinder-GetAttendanceByID] Error finding attendance by ID:", err)
		return nil, err
	}

	if attendance == nil {
		log.Println("[AttendanceFinder-GetAttendanceByID] Attendance not found for ID:", ID)
		return nil, errors.ErrInternalServerError.Error()
	}
	return attendance, nil
}

func (af *AttendanceFinder) GetAttendance(ctx context.Context, query, sort, order string, limit, page int) ([]*entity.Attendance, int64, error) {
	offset := (page - 1) * limit
	attendances, total, err := af.attendanceRepo.GetAttendance(ctx, query, sort, order, limit, offset)
	if err != nil {
		log.Fatal("[AttendanceFinder-GetAttendance] Error finding attendances:", err)
		return nil, 0, err
	}

	return attendances, total, nil
}

func (af *AttendanceFinder) GetAttendanceByUserID(ctx context.Context, userID uuid.UUID, query, sort, order string, limit, page int) ([]*entity.Attendance, int64, error) {
	offset := (page - 1) * limit
	attendances, total, err := af.attendanceRepo.GetAttendanceByUserID(ctx, userID, query, sort, order, limit, offset)
	if err != nil {
		log.Fatal("[AttendanceFinder-GetAttendanceByUserID] Error finding attendance by user ID:", err)
		return nil, 0, err
	}
	return attendances, total, nil
}
