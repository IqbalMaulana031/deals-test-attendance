package entity

import (
	"starter-go-gin/utils"
	"time"

	"github.com/google/uuid"
)

const (
	shiftDetailTableName = "master.shift_details"
)

// shiftDetail defines table shiftDetail
type ShiftDetail struct {
	ID          uuid.UUID `json:"id"`
	ShiftID     uuid.UUID `json:"shift_id"`
	Code        string    `json:"code"`
	DayType     bool      `json:"day_type"`
	DayInNumber int       `json:"day_in_number"`
	StartTime   string    `json:"start_time"`
	EndTime     string    `json:"end_time"`
	Auditable
}

// TableName specifies table name
func (model *ShiftDetail) TableName() string {
	return shiftDetailTableName
}

// NewShiftDetail creates new shiftDetail entity
func NewShiftDetail(
	id uuid.UUID,
	ShiftID uuid.UUID,
	Code string,
	DayType bool,
	DayInNumber int,
	StartTime string,
	EndTime string,
	createdBy string,
) *ShiftDetail {
	return &ShiftDetail{
		ID:          id,
		ShiftID:     ShiftID,
		Code:        Code,
		DayType:     DayType,
		DayInNumber: DayInNumber,
		StartTime:   StartTime,
		EndTime:     EndTime,
		Auditable:   NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *ShiftDetail) MapUpdateFrom(from *ShiftDetail) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"shift_id":      model.ShiftID,
			"code":          model.Code,
			"day_type":      model.DayType,
			"day_in_number": model.DayInNumber,
			"start_time":    model.StartTime,
			"end_time":      model.EndTime,
			"updated_by":    model.CreatedBy,
			"updated_at":    model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.ShiftID != from.ShiftID {
		mapped["shift_id"] = from.ShiftID
	}

	if model.Code != from.Code {
		mapped["code"] = from.Code
	}

	if model.DayType != from.DayType {
		mapped["day_type"] = from.DayType
	}

	if model.DayInNumber != from.DayInNumber {
		mapped["day_in_number"] = from.DayInNumber
	}

	if model.StartTime != from.StartTime {
		mapped["start_time"] = from.StartTime
	}

	if model.EndTime != from.EndTime {
		mapped["end_time"] = from.EndTime
	}

	mapped["updated_at"] = utils.AddSevenHours(time.Now())
	return &mapped
}
