package resource

import (
	"starter-go-gin/entity"

	"github.com/google/uuid"
)

type AttendanceRequest struct {
	ShiftDetailID uuid.UUID `json:"shift_detail_id" binding:"required,uuid"`
}

type OfferTimeRequest struct {
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time" binding:"required"`
}

type GetAttendanceListResponse struct {
	List  []*entity.Attendance `json:"list"`
	Total int64                `json:"total"`
	Meta  *Meta                `json:"meta"`
}
