package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	commonCache "starter-go-gin/common/cache"
	"starter-go-gin/common/constant"
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/logger"
	"starter-go-gin/common/tools"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/utils"
)

// ShiftRepository is a repository for auth
type ShiftRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
	cfg   config.Config
}

// ShiftRepositoryUseCase is a repository for auth
type ShiftRepositoryUseCase interface {
	// GetShift finds all shift by filter
	GetShift(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Shift, int64, error)
	// GetShiftByID finds an shift by id
	GetShiftByID(ctx context.Context, id uuid.UUID) (*entity.Shift, error)
	// GetSiftAndDetailsByID finds an shift and details by id
	GetShiftAndDetailsByID(ctx context.Context, id uuid.UUID) (*entity.Shift, error)
	// CreateShift create shift
	CreateShift(ctx context.Context, shift *entity.Shift) error
	// UpdateShift update shift
	UpdateShift(ctx context.Context, shift *entity.Shift, id uuid.UUID) error
	// DeleteShift update stock
	DeleteShift(ctx context.Context, id uuid.UUID) error
}

// NewShiftRepository returns a auth repository
func NewShiftRepository(
	db *gorm.DB,
	cache interfaces.Cacheable,
	cfg config.Config,
) *ShiftRepository {
	return &ShiftRepository{db, cache, cfg}
}

// GetShift is a function to get all shift
func (ur *ShiftRepository) GetShift(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Shift, int64, error) {
	var shift []*entity.Shift
	var total int64
	var gormDB = ur.db.
		WithContext(ctx).
		Model(&entity.Shift{})

	gormDB.Count(&total)

	gormDB = gormDB.Limit(limit).
		Offset(offset * limit)

	if query != "" {
		gormDB = gormDB.
			Where("shift_name ILIKE ?", "%"+query+"%")
	}

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", tools.EscapeSpecial(sort), tools.EscapeSpecial(order)))

	if err := gormDB.Find(&shift).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[ShiftRepository-GetShift] error when looking up all shift")
	}

	return shift, total, nil
}

// GetShiftByID finds an shift by id
func (ur *ShiftRepository) GetShiftByID(ctx context.Context, id uuid.UUID) (*entity.Shift, error) {
	result := &entity.Shift{}

	bytes, _ := ur.cache.Get(fmt.Sprintf(
		commonCache.ShiftByID, id.String()))

	if bytes != nil {
		if err := json.Unmarshal(bytes, &result); err != nil {
			return nil, err
		}
		return result, nil
	}

	if err := ur.db.
		WithContext(ctx).
		Where("id = ?", id).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logger.Error(ctx, errors.Wrap(err, "[ShiftRepository-GetShiftByID] id not found"))
		return nil, errors.Wrap(err, "[ShiftRepository-GetShiftByID] id not found")
	}

	if err := ur.cache.Set(fmt.Sprintf(commonCache.ShiftByID, id), &result, commonCache.OneMonth); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateShift create shift
func (ur *ShiftRepository) CreateShift(ctx context.Context, shift *entity.Shift) error {
	if err := ur.db.
		WithContext(ctx).
		Model(&entity.Shift{}).
		Create(shift).
		Error; err != nil {
		return errors.Wrap(err, "[ShiftRepository-CreateShift] error while creating shift")
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.ShiftByID, shift.ID)); err != nil {
		return err
	}

	return nil
}

// UpdateShift is a function to update shift
func (ur *ShiftRepository) UpdateShift(ctx context.Context, shift *entity.Shift, id uuid.UUID) error {
	oldTime := shift.UpdatedAt
	shift.UpdatedAt = time.Now()
	if err := ur.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.Shift)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, shift.ID).Error; err != nil {
				logger.Error(ctx, err)
				return errors.Wrap(err, "[ShiftRepository-UpdateShift] error while locking shift")
			}
			if err := tx.Model(&entity.Shift{}).
				Where(`id`, shift.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(shift)).Error; err != nil {
				logger.Error(ctx, err)
				return errors.Wrap(err, "[ShiftRepository-UpdateShift] error while update shift")
			}
			return nil
		}); err != nil {
		shift.UpdatedAt = oldTime
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.ShiftByID, shift.ID)); err != nil {
		return err
	}

	return nil
}

// DeleteShift is a function to delete shift
func (nc *ShiftRepository) DeleteShift(ctx context.Context, id uuid.UUID) error {
	if err := nc.db.WithContext(ctx).
		Model(&entity.Shift{}).
		Where(`id = ?`, id).
		Updates(
			map[string]interface{}{
				"updated_at": utils.AddSevenHours(time.Now()),
				"deleted_at": utils.AddSevenHours(time.Now()),
				"deleted_by": "system",
			}).Error; err != nil {
		return errors.Wrap(err, "[ShiftRepository-DeleteShift] error when updating user data")
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.ShiftByID, id)); err != nil {
		return err
	}
	return nil
}

// GetShiftAndDetailsByID finds an shift and details by id
func (ur *ShiftRepository) GetShiftAndDetailsByID(ctx context.Context, id uuid.UUID) (*entity.Shift, error) {
	result := &entity.Shift{}

	bytes, _ := ur.cache.Get(fmt.Sprintf(
		commonCache.ShiftAndDetailsByID, id.String()))

	if bytes != nil {
		if err := json.Unmarshal(bytes, &result); err != nil {
			return nil, err
		}
		return result, nil
	}

	if err := ur.db.
		WithContext(ctx).
		Preload("ShiftDetails").
		Where("id = ?", id).
		First(result).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logger.Error(ctx, errors.Wrap(err, "[ShiftRepository-GetShiftAndDetailsByID] id not found"))
		return nil, errors.Wrap(err, "[ShiftRepository-GetShiftAndDetailsByID] id not found")
	}

	if err := ur.cache.Set(fmt.Sprintf(commonCache.ShiftAndDetailsByID, id), &result, commonCache.OneMonth); err != nil {
		return nil, err
	}

	return result, nil
}
