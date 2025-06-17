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

// AttendanceRepository is a repository for attendance
type AttendanceRepository struct {
	db    *gorm.DB
	cache interfaces.Cacheable
	cfg   config.Config
}

// AttendanceRepositoryUseCase is a repository for attendance
type AttendanceRepositoryUseCase interface {
	// GetAttendance finds all Attendance by filter
	GetAttendance(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Attendance, int64, error)
	// GetAttendanceByID finds an Attendance by id
	GetAttendanceByID(ctx context.Context, id uuid.UUID) (*entity.Attendance, error)
	// CreateAttendance create Attendance
	CreateAttendance(ctx context.Context, Attendance *entity.Attendance) error
	// UpdateAttendance update Attendance
	UpdateAttendance(ctx context.Context, Attendance *entity.Attendance, id uuid.UUID) error
	// DeleteAttendance update stock
	DeleteAttendance(ctx context.Context, id uuid.UUID) error
	// GetAttendanceByUserID finds an attendance by user ID
	GetAttendanceByUserID(ctx context.Context, userID uuid.UUID, query, sort, order string, limit, page int) ([]*entity.Attendance, int64, error)
	// GetAttendanceByUserIDAndDate finds an attendance by user ID and date
	GetAttendanceByUserIDAndDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Attendance, error)
}

// NewAttendanceRepository returns a attendance repository
func NewAttendanceRepository(
	db *gorm.DB,
	cache interfaces.Cacheable,
	cfg config.Config,
) *AttendanceRepository {
	return &AttendanceRepository{db, cache, cfg}
}

// GetAttendance is a function to get all Attendance
func (ur *AttendanceRepository) GetAttendance(ctx context.Context, query, sort, order string, limit, offset int) ([]*entity.Attendance, int64, error) {
	var Attendance []*entity.Attendance
	var total int64
	var gormDB = ur.db.
		WithContext(ctx).
		Model(&entity.Attendance{})

	gormDB.Count(&total)

	gormDB = gormDB.Limit(limit).
		Offset(offset * limit)

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", tools.EscapeSpecial(sort), tools.EscapeSpecial(order)))

	if err := gormDB.Find(&Attendance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[AttendanceRepository-GetAttendance] error when looking up all Attendance")
	}

	return Attendance, total, nil
}

// GetAttendanceByID finds an Attendance by id
func (ur *AttendanceRepository) GetAttendanceByID(ctx context.Context, id uuid.UUID) (*entity.Attendance, error) {
	result := &entity.Attendance{}

	bytes, _ := ur.cache.Get(fmt.Sprintf(
		commonCache.AttendanceByID, id.String()))

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
		logger.Error(ctx, errors.Wrap(err, "[AttendanceRepository-GetAttendanceByID] id not found"))
		return nil, errors.Wrap(err, "[AttendanceRepository-GetAttendanceByID] id not found")
	}

	if err := ur.cache.Set(fmt.Sprintf(commonCache.AttendanceByID, id), &result, commonCache.OneMonth); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateAttendance create Attendance
func (ur *AttendanceRepository) CreateAttendance(ctx context.Context, Attendance *entity.Attendance) error {
	if err := ur.db.
		WithContext(ctx).
		Model(&entity.Attendance{}).
		Create(Attendance).
		Error; err != nil {
		return errors.Wrap(err, "[AttendanceRepository-CreateAttendance] error while creating Attendance")
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.AttendanceByID, Attendance.ID)); err != nil {
		return err
	}

	return nil
}

// UpdateAttendance is a function to update Attendance
func (ur *AttendanceRepository) UpdateAttendance(ctx context.Context, Attendance *entity.Attendance, id uuid.UUID) error {
	oldTime := Attendance.UpdatedAt
	Attendance.UpdatedAt = time.Now()
	if err := ur.db.
		WithContext(ctx).
		Transaction(func(tx *gorm.DB) error {
			sourceModel := new(entity.Attendance)
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Find(&sourceModel, Attendance.ID).Error; err != nil {
				logger.Error(ctx, err)
				return errors.Wrap(err, "[AttendanceRepository-UpdateAttendance] error while locking Attendance")
			}
			if err := tx.Model(&entity.Attendance{}).
				Where(`id`, Attendance.ID).
				UpdateColumns(sourceModel.MapUpdateFrom(Attendance)).Error; err != nil {
				logger.Error(ctx, err)
				return errors.Wrap(err, "[AttendanceRepository-UpdateAttendance] error while update Attendance")
			}
			return nil
		}); err != nil {
		Attendance.UpdatedAt = oldTime
	}

	if err := ur.cache.BulkRemove(fmt.Sprintf(commonCache.AttendanceByID, Attendance.ID)); err != nil {
		return err
	}

	return nil
}

// DeleteAttendance is a function to delete Attendance
func (nc *AttendanceRepository) DeleteAttendance(ctx context.Context, id uuid.UUID) error {
	if err := nc.db.WithContext(ctx).
		Model(&entity.Attendance{}).
		Where(`id = ?`, id).
		Updates(
			map[string]interface{}{
				"updated_at": utils.AddSevenHours(time.Now()),
				"deleted_at": utils.AddSevenHours(time.Now()),
				"deleted_by": "system",
			}).Error; err != nil {
		return errors.Wrap(err, "[AttendanceRepository-DeleteAttendance] error when updating user data")
	}

	if err := nc.cache.BulkRemove(fmt.Sprintf(commonCache.AttendanceByID, id)); err != nil {
		return err
	}
	return nil
}

// GetAttendanceByUserID finds an attendance by user ID
func (ur *AttendanceRepository) GetAttendanceByUserID(ctx context.Context, userID uuid.UUID, query, sort, order string, limit, page int) ([]*entity.Attendance, int64, error) {
	var Attendance []*entity.Attendance
	var total int64
	var gormDB = ur.db.
		WithContext(ctx).
		Model(&entity.Attendance{}).
		Where("employee_id = ?", userID)

	gormDB.Count(&total)

	gormDB = gormDB.Limit(limit).
		Offset(page * limit)

	if order != constant.Ascending && order != constant.Descending {
		order = constant.Descending
	}

	if sort == "" {
		sort = "created_at"
	}

	gormDB = gormDB.Order(fmt.Sprintf("%s %s", tools.EscapeSpecial(sort), tools.EscapeSpecial(order)))

	if err := gormDB.Find(&Attendance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, 0, errors.Wrap(err, "[AttendanceRepository-GetAttendanceByUserID] error when looking up all Attendance by user ID")
	}

	return Attendance, total, nil
}

// GetAttendanceByUserIDAndDate finds an attendance by user ID and date
func (ur *AttendanceRepository) GetAttendanceByUserIDAndDate(ctx context.Context, userID uuid.UUID, date time.Time) (*entity.Attendance, error) {
	var attendance entity.Attendance

	if err := ur.db.
		WithContext(ctx).
		Where("employee_id = ? AND DATE(attendance_date) = ?", userID, date.Format("2006-01-02")).
		First(&attendance).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		logger.Error(ctx, errors.Wrap(err, "[AttendanceRepository-GetAttendanceByUserIDAndDate] user ID and date not found"))
		return nil, errors.Wrap(err, "[AttendanceRepository-GetAttendanceByUserIDAndDate] user ID and date not found")
	}

	return &attendance, nil
}
