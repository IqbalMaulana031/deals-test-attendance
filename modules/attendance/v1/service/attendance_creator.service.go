package service

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"starter-go-gin/entity"
	"time"

	"github.com/google/uuid"
)

func (ac *AttendanceCreator) RechordTime(ctx context.Context, userID, shiftDetailID uuid.UUID) error {
	today := time.Now()
	// find shift by ID
	shift, err := ac.ShiftDetailRepo.GetShiftDetailByID(ctx, shiftDetailID)
	if err != nil {
		log.Println("[AttendanceCreator-RechordTime] Error finding shift detail by ID:", err)
		return errors.New("error finding shift")
	}
	if shift == nil {
		log.Println("[AttendanceCreator-RechordTime] Shift detail not found for ID:", shiftDetailID)
		return errors.New("shift not found")
	}
	// check if attendance at today already exists
	exists, err := ac.AttendanceRepository.GetAttendanceByUserIDAndDate(ctx, userID, today)
	if err != nil {
		log.Println("[AttendanceCreator-RechordTime] Error checking attendance:", err)
		return errors.New("error checking attendance")
	}
	if exists == nil {
		err := ac.AttendanceRepository.CreateAttendance(ctx, &entity.Attendance{
			ID:             uuid.New(),
			EmployeeID:     userID,
			ShiftID:        shiftDetailID,
			ShiftStart:     shift.StartTime,
			ShiftEnd:       shift.EndTime,
			AttendanceDate: today,
			Checkin:        today,
			Auditable: entity.Auditable{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				CreatedBy: sql.NullString{String: userID.String(), Valid: true},
				UpdatedBy: sql.NullString{String: userID.String(), Valid: true},
			},
		})
		if err != nil {
			log.Println("[AttendanceCreator-RechordTime] Error creating attendance:", err)
			return errors.New("error creating attendance")
		}
	} else {
		// Update existing attendance
		exists.Checkout = today
		exists.UpdatedAt = time.Now()
		exists.UpdatedBy = sql.NullString{String: userID.String(), Valid: true}

		err := ac.AttendanceRepository.UpdateAttendance(ctx, exists, exists.ID)
		if err != nil {
			log.Println("[AttendanceCreator-RechordTime] Error updating attendance:", err)
			return errors.New("error updating attendance")
		}
		log.Println("[AttendanceCreator-RechordTime] Attendance updated successfully")
		log.Printf("User %s checked in at %s for shift %s", userID, today, shiftDetailID)
		log.Printf("Shift details: Start - %s, End - %s", shift.StartTime, shift.EndTime)
	}
	return nil
}
