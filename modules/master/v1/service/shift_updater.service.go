package service

import (
	"context"
	"errors"
	"log"
	"starter-go-gin/entity"
)

func (mu *MasterUpdater) UpdateShift(ctx context.Context, shift *entity.Shift) error {
	err := mu.shiftRepo.UpdateShift(ctx, shift, shift.ID)
	if err != nil {
		log.Println("[MasterUpdater-UpdateShift] Error updating shift:", err)
		return errors.New("failed to update shift")
	}
	return nil
}
