package resource

import "github.com/google/uuid"

// create shift detail request
type CreateShiftDetailRequest struct {
	ShiftID     uuid.UUID `json:"shift_id"`
	Code        string    `json:"code"`
	DayType     bool      `json:"day_type"`
	DayInNumber int       `json:"day_in_number"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
}

type ShiftDetail struct {
	ID          uuid.UUID `json:"id"`
	ShiftID     uuid.UUID `json:"shift_id"`
	Code        string    `json:"code"`
	DayType     bool      `json:"day_type"`
	DayInNumber int       `json:"day_in_number"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
}
