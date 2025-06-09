package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/entity"
)

// CreateUserRole creates a new user role
func (uc *UserCreator) CreateUserRole(ctx context.Context, userID uuid.UUID, roleID uuid.UUID) (*entity.UserRole, error) {
	userRole := entity.NewUserRole(
		uuid.New(),
		userID,
		roleID,
		"system")

	err := uc.userRoleRepo.CreateOrUpdate(ctx, userRole)
	if err != nil {
		return nil, err
	}

	return userRole, nil
}
