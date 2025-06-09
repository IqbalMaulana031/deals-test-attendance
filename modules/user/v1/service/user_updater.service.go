package service

import (
	"context"

	"starter-go-gin/common/errors"
	"starter-go-gin/entity"
)

// UpdateUser updates a user
func (uu *UserUpdater) UpdateUser(ctx context.Context, user *entity.User) error {
	if err := uu.userRepo.Update(ctx, user); err != nil {
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
