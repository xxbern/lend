package logx

import (
	"context"
	"time"

	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type GormLogger struct {
	zlog    *zap.Logger
	level   logger.LogLevel
	slowSQL time.Duration
}

func NewGormLogger(level logger.LogLevel, slow time.Duration) logger.Interface {
	return &GormLogger{
		zlog:    Logger,
		level:   level,
		slowSQL: slow,
	}
}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Info {
		l.zlog.Sugar().Infof(msg, data...)
	}
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Warn {
		l.zlog.Sugar().Warnf(msg, data...)
	}
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.level >= logger.Error {
		l.zlog.Sugar().Errorf(msg, data...)
	}
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	fields := []zap.Field{
		zap.Duration("elapsed", elapsed),
		zap.Int64("rows", rows),
		zap.String("sql", sql),
	}

	switch {
	case err != nil && l.level >= logger.Error:
		l.zlog.Error("gorm error", append(fields, zap.Error(err))...)
	case elapsed > l.slowSQL && l.level >= logger.Warn:
		l.zlog.Warn("gorm slow query", fields...)
	case l.level >= logger.Info:
		l.zlog.Info("gorm query", fields...)
	}
}
