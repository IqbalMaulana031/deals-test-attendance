package resource

import (
	"starter-go-gin/entity"

	"github.com/google/uuid"
)

type CreateShiftRequest struct {
	ShiftName  string                     `json:"shift_name" binding:"required,max=50"`
	ShifDetail []CreateShiftDetailRequest `json:"shift_detail" binding:"required"`
}

type UpdateShiftRequest struct {
	ID        uuid.UUID `json:"id" binding:"required,uuid"`
	ShiftName string    `json:"shift_name" binding:"required,max=50"`
	IsDefault bool      `json:"is_default"`
}

type GetShiftByIDRequest struct {
	ID uuid.UUID `uri:"id" binding:"required"`
}

type Shift struct {
	ID        uuid.UUID `json:"id"`
	ShiftName string    `json:"shift_name"`
	IsDefault bool      `json:"is_default"`
}

type GetShiftListResponse struct {
	List  []*entity.Shift `json:"list"`
	Total int64           `json:"total"`
	Meta  *Meta           `json:"meta"`
}

type GetShiftAndDetailsByIDResponse struct {
	ID        uuid.UUID             `json:"id"`
	ShiftName string                `json:"shift_name"`
	IsDefault bool                  `json:"is_default"`
	Details   []*entity.ShiftDetail `json:"details"`
}
