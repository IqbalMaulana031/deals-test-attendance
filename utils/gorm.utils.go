package utils

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"starter-go-gin/common/logger"
	"starter-go-gin/config"
)

// NewPostgresGormDB builds a connection of gorm to PostgreSQL.
func NewPostgresGormDB(cfg *config.Config, log *logger.GormLogger) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Name,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn))
	if log != nil {
		gormDB.Logger = log
	} else {
		gormDB.Logger = gormLogger.Default.LogMode(gormLogger.Info)
	}

	sqlDB, sdErr := gormDB.DB()
	if sdErr != nil {
		return nil, sdErr
	}

	maxOpenConns, mocErr := strconv.Atoi(cfg.Postgres.MaxOpenConns)
	if mocErr != nil {
		return nil, mocErr
	}

	maxConnLifetime, mclErr := time.ParseDuration(cfg.Postgres.MaxConnLifetime)
	if mclErr != nil {
		return nil, mclErr
	}

	maxIdleLifetime, milErr := time.ParseDuration(cfg.Postgres.MaxIdleLifetime)
	if milErr != nil {
		return nil, milErr
	}

	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(maxConnLifetime)
	sqlDB.SetConnMaxIdleTime(maxIdleLifetime)

	return gormDB, err
}
