package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/entity"
)

// FindUserRoleByUserID is a function that finds user role by user id
func (uf *UserFinder) FindUserRoleByUserID(ctx context.Context, userID uuid.UUID) (*entity.UserRole, error) {
	userRole, err := uf.userRoleRepo.FindByUserID(ctx, userID)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	if userRole == nil {
		return nil, errors.ErrAccountIsBroken.Error()
	}

	return userRole, nil
}
