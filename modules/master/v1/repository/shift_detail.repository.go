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
	"starter-go-gin/common/interfaces"
	"starter-go-gin/common/logger"
	"starter-go-gin/config"
	"starter-go-gin/entity"
	"starter-go-gin/utils"
)

// ShiftDetailRepository is a repository for auth
type ShiftDetailRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
	cfg   config.Config
}

// ShiftDetailRepositoryUseCase is a repository for auth
type ShiftDetailRepositoryUseCase interface {
	// GetShiftDetailByID finds an shift detail by id
	GetShiftDetailByID(ctx context.Context, id uuid.UUID) (*entity.ShiftDetail, error)
	// CreateShiftDetail create shift detail
	CreateShiftDetail(ctx context.Context, shiftDetail *entity.ShiftDetail) error
	// UpdateShiftDetail update shift detail
	UpdateShiftDetail(ctx context.Context, shiftDetail *entity.ShiftDetail, id uuid.UUID) error
	// DeleteShiftDetail delete shift detail
	DeleteShiftDetail(ctx context.Context, id uuid.UUID) error
}

// NewShiftDetailRepository returns a auth repository
func NewShiftDetailRepository(
	db *gorm.DB,
	cache interfaces.Cacheable,
	cfg config.Config,
) *ShiftDetailRepository {
	return &ShiftDetailRepository{db, cache, cfg}
}

// GetShiftDetailByID is a function to get shift detail by id
func (ur *ShiftDetailRepository) GetShiftDetailByID(ctx context.Context, id uuid.UUID) (*entity.ShiftDetail, error) {
	result := &entity.ShiftDetail{}

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
		logger.Error(ctx, errors.Wrap(err, "[ShiftDetailRepository-GetShiftDetailByID] id not found"))
		return nil, errors.Wrap(err, "[ShiftDetailRepository-GetShiftDetailByID] id not found")
	}

	if err := ur.cache.Set(fmt.Sprintf(commonCache.ShiftByID, id), &result, commonCache.OneMonth); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateShiftDetail create shift detail
func (ur *ShiftDetailRepository) CreateShiftDetail(ctx context.Context, shiftDetail *entity.ShiftDetail) error {
	if shiftDetail == nil {
		return errors.New("shift detail is nil")
	}

	if err := ur.db.WithContext(ctx).Create(shiftDetail).Error; err != nil {
		return errors.Wrap(err, "[ShiftDetailRepository-CreateShiftDetail] error while creating shift detail")
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.ShiftDetailByID, shiftDetail.ID)); err != nil {
		return err
	}

	return nil
}

// UpdateShiftDetail update shift detail
func (ur *ShiftDetailRepository) UpdateShiftDetail(ctx context.Context, shiftDetail *entity.ShiftDetail, id uuid.UUID) error {
	oldTime := shiftDetail.UpdatedAt
	shiftDetail.UpdatedAt = time.Now()
	if err := ur.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.ShiftDetail)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, shiftDetail.ID).Error; err != nil {
				logger.Error(ctx, err)
				return errors.Wrap(err, "[ShiftRepository-UpdateShift] error while locking shiftDetail")
			}
			if err := tx.Model(&entity.ShiftDetail{}).
				Where(`id`, shiftDetail.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(shiftDetail)).Error; err != nil {
				logger.Error(ctx, err)
				return errors.Wrap(err, "[ShiftRepository-UpdateShift] error while update shiftDetail")
			}
			return nil
		}); err != nil {
		shiftDetail.UpdatedAt = oldTime
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.ShiftByID, shiftDetail.ID)); err != nil {
		return err
	}

	return nil
}

// DeleteShiftDetail delete shift detail
func (ur *ShiftDetailRepository) DeleteShiftDetail(ctx context.Context, id uuid.UUID) error {
	if err := ur.db.WithContext(ctx).
		Model(&entity.ShiftDetail{}).
		Where(`id = ?`, id).
		Updates(
			map[string]interface{}{
				"updated_at": utils.AddSevenHours(time.Now()),
				"deleted_at": utils.AddSevenHours(time.Now()),
				"deleted_by": "system",
			}).Error; err != nil {
		return errors.Wrap(err, "[ShiftDetailRepository-DeleteShiftDetail] error when updating shift detail data")
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.ShiftByID, id)); err != nil {
		return err
	}
	return nil
}
