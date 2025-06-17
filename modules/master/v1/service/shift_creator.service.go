package service

import (
	"context"
	"errors"
	"log"
	"starter-go-gin/entity"
)

func (mc *MasterCreator) CreateShift(ctx context.Context, shift *entity.Shift, shiftDetail *[]entity.ShiftDetail) error {
	// create a new shift in the repository
	err := mc.shiftRepo.CreateShift(ctx, shift)
	if err != nil {
		log.Fatal("[MasterCreator-CreateShift] Error creating shift:", err)
		return errors.New("failed to create shift")
	}

	// create shift details if create shift is successful
	if shiftDetail != nil {
		for _, detail := range *shiftDetail {
			detail.ShiftID = shift.ID // associate the detail with the created shift
			err = mc.shiftDetailRepo.CreateShiftDetail(ctx, &detail)
			if err != nil {
				log.Fatal("[MasterCreator-CreateShift] Error creating shift detail:", err)
				return errors.New("failed to create shift detail")
			}
		}
	}
	return nil
}
