package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/entity"
)

// GetUsers gets all users
func (uf *UserFinder) GetUsers(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.User, int64, error) {
	users, total, err := uf.userRepo.GetUsers(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, 0, errors.ErrInternalServerError.Error()
	}

	return users, total, nil
}

// GetUserByID gets user by id
func (uf *UserFinder) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := uf.userRepo.GetUserByID(ctx, id)

	if err != nil {
		return user, errors.ErrInternalServerError.Error()
	}

	if user == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}

	return user, nil
}

// GetUserByUsername gets user by phone number
func (uf *UserFinder) GetUserByUsername(ctx context.Context, username, roleName string) (*entity.User, error) {
	user, err := uf.userRepo.GetUserByUsername(ctx, username, roleName)

	if err != nil {
		return user, errors.ErrInternalServerError.Error()
	}

	if user == nil {
		return nil, errors.ErrRecordNotFound.Error()
	}
	return user, nil
}
