package service

import (
	"context"

	"starter-go-gin/common/errors"
	"starter-go-gin/entity"
)

// CreateUser creates a new user
func (uc *UserCreator) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	err := uc.userRepo.CreateUser(ctx, user)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return user, nil
}
