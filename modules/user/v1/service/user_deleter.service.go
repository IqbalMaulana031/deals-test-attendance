package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
)

// DeleteUser deletes employee
func (ud *UserDeleter) DeleteUser(ctx context.Context, id, merchantID uuid.UUID) error {
	if err := ud.userRepo.DeleteUser(ctx, id); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
