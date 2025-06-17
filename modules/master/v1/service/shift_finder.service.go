package service

import (
	"context"
	"log"
	"starter-go-gin/common/errors"
	"starter-go-gin/entity"

	"github.com/google/uuid"
)

func (mf *MasterFinder) GetShiftByID(ctx context.Context, ID uuid.UUID) (*entity.Shift, error) {
	shift, err := mf.shiftRepo.GetShiftByID(ctx, ID)
	if err != nil {
		log.Fatal("[ShiftFinder-GetShiftByID] Error finding shift by ID:", err)
		return nil, err
	}

	if shift == nil {
		log.Println("[ShiftFinder-GetShiftByID] Shift not found for ID:", ID)
		return nil, errors.ErrInternalServerError.Error()
	}

	return shift, nil
}

func (mf *MasterFinder) GetShift(ctx context.Context, query, sort, order string, limit, page int) ([]*entity.Shift, int64, error) {
	offset := (page - 1) * limit
	shifts, total, err := mf.shiftRepo.GetShift(ctx, query, sort, order, limit, offset)
	if err != nil {
		log.Fatal("[ShiftFinder-GetShift] Error finding shifts:", err)
		return nil, 0, err
	}

	if shifts == nil {
		log.Println("[ShiftFinder-GetShift] No shifts found")
		return nil, 0, errors.ErrInternalServerError.Error()
	}

	return shifts, total, nil
}

func (mf *MasterFinder) GetShiftAndDetailsByID(ctx context.Context, ID uuid.UUID) (*entity.Shift, error) {
	shift, err := mf.shiftRepo.GetShiftAndDetailsByID(ctx, ID)
	if err != nil {
		log.Fatal("[ShiftFinder-GetShiftAndDetailsByID] Error finding shift and details by ID:", err)
		return nil, err
	}

	if shift == nil {
		log.Println("[ShiftFinder-GetShiftAndDetailsByID] Shift not found for ID:", ID)
		return nil, errors.ErrInternalServerError.Error()
	}

	return shift, nil
}
