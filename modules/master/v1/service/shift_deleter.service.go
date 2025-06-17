package service

import (
	"context"
	"errors"
	"log"

	"github.com/google/uuid"
)

func (md *MasterDeleter) DeleteShift(ctx context.Context, id uuid.UUID) error {
	// delete the shift itself
	if err := md.shiftRepo.DeleteShift(ctx, id); err != nil {
		log.Println("[MasterDeleter-DeleteShift] Error deleting shift:", err)
		return errors.New("failed to delete shift")
	}

	return nil
}
