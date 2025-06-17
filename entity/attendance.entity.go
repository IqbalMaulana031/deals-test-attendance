package entity

import (
	"time"

	"github.com/google/uuid"
)

const (
	attendanceTableName = "attendance.attendances"
)

// Attendance defines table attendance
type Attendance struct {
	ID             uuid.UUID `json:"id"`
	EmployeeID     uuid.UUID `json:"employee_id"`
	ShiftID        uuid.UUID `json:"shift_id"`
	ShiftStart     string    `json:"shift_start"`
	ShiftEnd       string    `json:"shift_end"`
	AttendanceDate time.Time `json:"attendance_date"`
	Checkin        time.Time `json:"checkin"`
	Checkout       time.Time `json:"checkout"`
	Auditable
}

// TableName returns the table name for the Attendance entity
func (Attendance) TableName() string {
	return attendanceTableName
}

// newAttendance creates a new Attendance entity
func NewAttendance(
	id uuid.UUID,
	employeeID uuid.UUID,
	shiftID uuid.UUID,
	shiftStart string,
	shiftEnd string,
	attendanceDate time.Time,
	checkIn time.Time,
	checkOut time.Time,
	createdBy string,
) *Attendance {
	return &Attendance{
		ID:             id,
		EmployeeID:     employeeID,
		ShiftID:        shiftID,
		ShiftStart:     shiftStart,
		ShiftEnd:       shiftEnd,
		AttendanceDate: attendanceDate,
		Checkin:        checkIn,
		Checkout:       checkOut,
		Auditable:      NewAuditable(createdBy),
	}
}

// MapUpdateFrom maps fields from another Attendance entity for updates
func (model *Attendance) MapUpdateFrom(from *Attendance) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"employee_id":     model.EmployeeID,
			"shift_id":        model.ShiftID,
			"shift_start":     model.ShiftStart,
			"shift_end":       model.ShiftEnd,
			"attendance_date": model.AttendanceDate,
			"checkin":         model.Checkin,
			"checkout":        model.Checkout,
			"updated_by":      model.CreatedBy,
			"updated_at":      model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.EmployeeID != from.EmployeeID {
		mapped["employee_id"] = from.EmployeeID
	}
	if model.ShiftID != from.ShiftID {
		mapped["shift_id"] = from.ShiftID
	}
	if model.ShiftStart != from.ShiftStart {
		mapped["shift_start"] = from.ShiftStart
	}
	if model.ShiftEnd != from.ShiftEnd {
		mapped["shift_end"] = from.ShiftEnd
	}
	if !model.AttendanceDate.Equal(from.AttendanceDate) {
		mapped["attendance_date"] = from.AttendanceDate
	}
	if !model.Checkin.Equal(from.Checkin) {
		mapped["checkin"] = from.Checkin
	}
	if !model.Checkout.Equal(from.Checkout) {
		mapped["checkout"] = from.Checkout
	}

	mapped["updated_at"] = time.Now()
	return &mapped
}
