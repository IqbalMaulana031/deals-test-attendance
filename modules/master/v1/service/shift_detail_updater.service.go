package service

import (
	"context"
	"errors"
	"log"
	"starter-go-gin/entity"
)

// UpdateShiftDetail updates an existing shift detail
func (mu *MasterUpdater) UpdateShiftDetail(ctx context.Context, shiftDetail *entity.ShiftDetail) error {
	err := mu.shiftDetailRepo.UpdateShiftDetail(ctx, shiftDetail, shiftDetail.ID)
	if err != nil {
		log.Println("[MasterUpdater-UpdateShiftDetail] Error updating shift detail:", err)
		return errors.New("failed to update shift detail")
	}
	return nil
}
