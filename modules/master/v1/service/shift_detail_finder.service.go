package service

import (
	"context"
	"log"
	"net/http"
	"starter-go-gin/common/errors"
	"starter-go-gin/entity"

	"github.com/google/uuid"
)

func (mf *MasterFinder) GetShiftDetailByID(ctx context.Context, ID uuid.UUID) (*entity.ShiftDetail, error) {
	shiftDetail, err := mf.shiftDetailRepo.GetShiftDetailByID(ctx, ID)
	if err != nil {
		log.Fatal("[ShiftDetailFinder-GetShiftDetailByID] Error finding shift detail by ID:", err)
		return nil, errors.NewError(http.StatusBadRequest, "Error finding shift detail by ID").Error()
	}

	if shiftDetail == nil {
		log.Println("[ShiftDetailFinder-GetShiftDetailByID] Shift detail not found for ID:", ID)
		return nil, errors.ErrRecordNotFound.Error()
	}

	return shiftDetail, nil
}
