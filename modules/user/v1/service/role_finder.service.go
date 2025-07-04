package service

import (
	"context"

	"github.com/google/uuid"

	"starter-go-gin/common/errors"
	"starter-go-gin/entity"
)

// GetRoleByName is a function that gets role by name
func (uf *UserFinder) GetRoleByName(ctx context.Context, name string) (*entity.Role, error) {
	role, err := uf.roleRepo.FindByName(ctx, name)
	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return role, nil
}

// GetRoleByID is a function that gets role by id
func (uf *UserFinder) GetRoleByID(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	role, err := uf.roleRepo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return role, nil
}

// GetRoles gets all roles
func (uf *UserFinder) GetRoles(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Role, error) {
	roles, err := uf.roleRepo.FindAll(ctx, query, sort, order, limit, offset)

	if err != nil {
		return nil, errors.ErrInternalServerError.Error()
	}

	return roles, nil
}
