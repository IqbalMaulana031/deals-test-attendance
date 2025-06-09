package logger

import (
	"context"
	"strings"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	logLevel gormlogger.LogLevel
}

func NewGormLogger(logLevel gormlogger.LogLevel) *GormLogger {
	return &GormLogger{logLevel: logLevel}
}

func (l *GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	l.logLevel = level
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= gormlogger.Info {
		Info(ctx, "[INFO] "+msg, data...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= gormlogger.Warn {
		Warn(ctx, "[WARN] "+msg, data...)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= gormlogger.Error {
		// loop and ignore if data contains string record not found
		for _, d := range data {
			// check if d is string and contains record not found
			if str, ok := d.(string); ok && strings.Contains(str, "record not found") {
				return
			}
		}
		ErrorFromStr(ctx, data...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.logLevel > gormlogger.Silent {
		elapsed := time.Since(begin)
		sql, rows := fc()
		if err != nil {
			l.Error(ctx, "trace error: %v | sql: %s | rows: %d | elapsed: %s", err, sql, rows, elapsed)
		} else {
			l.Info(ctx, "trace success | sql: %s | rows: %d | elapsed: %s", sql, rows, elapsed)
		}
	}
}
