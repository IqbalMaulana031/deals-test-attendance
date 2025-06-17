package entity

import (
	"starter-go-gin/utils"
	"time"

	"github.com/google/uuid"
)

const (
	shiftTableName = "master.shifts"
)

// shift defines table shift
type Shift struct {
	ID           uuid.UUID      `json:"id"`
	ShiftName    string         `json:"shift_name"`
	IsDefault    bool           `json:"is_default"`
	ShiftDetails []*ShiftDetail `json:"shift_details,omitempty"`
	Auditable
}

// TableName specifies table name
func (model *Shift) TableName() string {
	return shiftTableName
}

// NewShift creates new shift entity
func NewShift(
	id uuid.UUID,
	shiftName string,
	createdBy string,
) *Shift {
	return &Shift{
		ID:        id,
		ShiftName: shiftName,
		Auditable: NewAuditable(createdBy),
	}
}

// MapUpdateFrom mapping from model
func (model *Shift) MapUpdateFrom(from *Shift) *map[string]interface{} {
	if from == nil {
		return &map[string]interface{}{
			"shift_name": model.ShiftName,
			"updated_by": model.CreatedBy,
			"updated_at": model.UpdatedAt,
		}
	}

	mapped := make(map[string]interface{})

	if model.ShiftName != from.ShiftName {
		mapped["shift_name"] = from.ShiftName
	}

	mapped["updated_at"] = utils.AddSevenHours(time.Now())
	return &mapped
}
