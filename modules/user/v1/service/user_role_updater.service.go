package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/common/logger"
)

// UpdateUserRole updates a user role
func (uu *UserUpdater) UpdateUserRole(ctx context.Context, userID, roleID uuid.UUID) error {
	userRole, err := uu.userRoleRepo.FindByUserID(ctx, userID)

	if err != nil {
		logger.ErrorWithStr(ctx, "Error when find user role: ", err)
		return errors.ErrInternalServerError.Error()
	}

	if userRole == nil {
		logger.ErrorFromStr(ctx, "user role is nil")
		return errors.ErrRecordNotFound.Error()
	}

	userRole.RoleID = roleID

	if err := uu.userRoleRepo.Update(ctx, userRole); err != nil {
		logger.ErrorWithStr(ctx, "Error when update user role: ", err)
		return errors.ErrInternalServerError.Error()
	}

	return nil
}
