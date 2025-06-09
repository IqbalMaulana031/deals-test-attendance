package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	commonCache "starter-go-gin/common/cache"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/tools"
	"starter-go-gin/entity"
)

// RoleRepository is a repository for role
type RoleRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
}

// RoleRepositoryUseCase is a use case for role
type RoleRepositoryUseCase interface {
	// Create creates a role
	Create(ctx context.Context, role *entity.Role) error
	// FindAll finds all roles
	FindAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Role, error)
	// FindByID finds a role by id
	FindByID(ctx context.Context, id uuid.UUID) (*entity.Role, error)
	// Delete deletes a role
	Delete(ctx context.Context, id uuid.UUID, deletedBy string) error
	// FindByName finds a role by name
	FindByName(ctx context.Context, slug string) (*entity.Role, error)
	// Update update a role
	Update(ctx context.Context, role *entity.Role) error
	// FindByType finds a role by type
	FindByType(ctx context.Context, roleType, query string) ([]*entity.Role, error)
}

// NewRoleRepository creates a new role repository
func NewRoleRepository(db *gorm.DB, cache interfaces.Cacheable) *RoleRepository {
	return &RoleRepository{db, cache}
}

// Create creates a role
func (nc *RoleRepository) Create(ctx context.Context, role *entity.Role) error {
	if err := nc.db.WithContext(ctx).Model(&entity.Role{}).Create(role).Error; err != nil {
		return errors.Wrap(err, "[RoleRepository-CreateRole] error while creating role")
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.RoleFindByID, "*")); err != nil {
		return err
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.RoleFindByType, "*", "*")); err != nil {
		return err
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.UserRoleFindByRoleID, '*')); err != nil {
		return err
	}

	return nil
}

// FindByID finds a role by id
func (nc *RoleRepository) FindByID(ctx context.Context, id uuid.UUID) (*entity.Role, error) {
	role := &entity.Role{}

	bytes, _ := nc.cache.Get(fmt.Sprintf(
		commonCache.RoleFindByID,
		id,
	))

	if bytes != nil {
		if err := json.Unmarshal(bytes, &role); err != nil {
			return nil, err
		}
		return role, nil
	}

	if err := nc.db.
		WithContext(ctx).
		Model(&entity.Role{}).
		Where("id = ?", id).
		First(&role).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[RoleRepository-FindByID] error while getting role")
	}

	if err := nc.cache.Set(fmt.Sprintf(
		commonCache.RoleFindByID,
		id,
	), role, commonCache.OneMonth); err != nil {
		return nil, err
	}

	return role, nil
}

// FindAll finds all roles
func (nc *RoleRepository) FindAll(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Role, error) {
	role := make([]*entity.Role, 0)
	var gormDB = nc.db.
		WithContext(ctx).
		Model(&entity.Role{})

	if query != "" {
		gormDB.Where("name ILIKE ?", "%"+query+"%")
	}

	if sort != "" {
		gormDB.Order(fmt.Sprintf("%s %s", tools.EscapeSpecial(sort), tools.EscapeSpecial(order)))
	}

	if limit > 0 {
		gormDB.Limit(limit)
	}

	if offset > 0 {
		gormDB.Offset(offset)
	}

	if err := gormDB.
		Find(&role).
		Error; err != nil {
		return nil, errors.Wrap(err, "[RoleRepository-GetNewsCategories] error while getting news category")
	}

	return role, nil
}

// Delete deletes a role
func (nc *RoleRepository) Delete(ctx context.Context, id uuid.UUID, deletedBy string) error {
	if err := nc.db.WithContext(ctx).
		Model(&entity.Role{}).
		Where(`id = ?`, id).
		Updates(
			map[string]interface{}{
				"deleted_by": deletedBy,
				"updated_at": time.Now(),
				"deleted_at": time.Now(),
			}).Error; err != nil {
		return errors.Wrap(err, "[RoleRepository-DeactivateRole] error when updating role data")
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.RoleFindByID, "*")); err != nil {
		return err
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.RoleFindByType, "*", "*")); err != nil {
		return err
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.UserRoleFindByRoleID, '*')); err != nil {
		return err
	}
	return nil
}

// FindByName finds a role by name
func (nc *RoleRepository) FindByName(ctx context.Context, name string) (*entity.Role, error) {
	role := &entity.Role{}

	if err := nc.db.
		WithContext(ctx).
		Model(&entity.Role{}).
		Where("LOWER(name) = ?", strings.ToLower(name)).
		First(&role).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(err, "[RoleRepository-FindByName] error while getting role")
	}

	return role, nil
}

// Update update a role
func (nc *RoleRepository) Update(ctx context.Context, role *entity.Role) error {
	role.UpdatedAt = time.Now()
	sourceModelNews := new(entity.Role)
	if err := nc.db.WithContext(ctx).Model(&entity.Role{}).
		Where("id = ?", role.ID).
		UpdateColumns(sourceModelNews.MapUpdateFrom(role)).
		Error; err != nil {
		return errors.Wrap(err, "[RoleRepository-Update] error while updating role")
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.RoleFindByID, "*")); err != nil {
		return err
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.RoleFindByType, "*", "*")); err != nil {
		return err
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.UserRoleFindByRoleID, '*')); err != nil {
		return err
	}

	return nil
}

// FindByType finds a role by type
func (nc *RoleRepository) FindByType(ctx context.Context, roleType, query string) ([]*entity.Role, error) {
	role := make([]*entity.Role, 0)

	bytes, _ := nc.cache.Get(fmt.Sprintf(
		commonCache.RoleFindByType,
		roleType, query,
	))

	if bytes != nil {
		if err := json.Unmarshal(bytes, &role); err != nil {
			return nil, err
		}
		return role, nil
	}

	gormDB := nc.db.
		WithContext(ctx).
		Model(&entity.Role{}).
		Where("type = ?", roleType)

	if query != "" {
		gormDB = gormDB.Where("name ILIKE ?", "%"+query+"%")
	}

	if err := gormDB.
		Find(&role).
		Error; err != nil {
		return nil, errors.Wrap(err, "[RoleRepository-FindByType] error while getting role")
	}

	if err := nc.cache.Set(fmt.Sprintf(commonCache.RoleFindByType, roleType, query), &role, commonCache.OneWeek); err != nil {
		return nil, err
	}

	return role, nil
}
