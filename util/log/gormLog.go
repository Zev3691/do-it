package log

import (
	"context"
	"fmt"
	"runtime"
	"strconv"
	"time"

	"gorm.io/gorm/logger"
)

type gormLogger struct {
	LogLevel      logger.LogLevel
	SlowThreshold time.Duration
}

func NewGormLogger() *gormLogger {
	return &gormLogger{
		LogLevel:      logger.Info,
		SlowThreshold: 10 * time.Second,
	}
}

// LogMode log mode
func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

// Info print info
func (l *gormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		FromContext(ctx).Sugar().Infof(msg, append([]interface{}{fileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l *gormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		FromContext(ctx).Sugar().Warnf(msg, append([]interface{}{fileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l *gormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		FromContext(ctx).Sugar().Errorf(msg, append([]interface{}{fileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel > 0 {
		elapsed := time.Since(begin)
		switch {
		case err != nil && l.LogLevel >= logger.Error:
			sql, rows := fc()
			if rows == -1 {
				FromContext(ctx).Sugar().Errorf("[err: %s, cost: %f, rows: %s, sql: %s]", err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				FromContext(ctx).Sugar().Errorf("[err: %s, cost: %f, rows: %d, sql: %s]", err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
			sql, rows := fc()
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			if rows == -1 {
				FromContext(ctx).Sugar().Warnf("[slowLog: %s, cost: %f, rows: %s, sql: %s]", slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				FromContext(ctx).Sugar().Warnf("[slowLog: %s, cost: %f, rows: %d, sql: %s]", slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		case l.LogLevel >= logger.Info:
			sql, rows := fc()
			if rows == -1 {
				FromContext(ctx).Sugar().Infof("[cost: %f, rows: %s, sql: %s]", float64(elapsed.Nanoseconds())/1e6, "-", sql)
			} else {
				FromContext(ctx).Sugar().Infof("[cost: %f, rows: %d, sql: %s]", float64(elapsed.Nanoseconds())/1e6, rows, sql)
			}
		}
	}
}

func fileWithLineNum() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)

		if ok {
			return file + ":" + strconv.FormatInt(int64(line), 10)
		}
	}
	return ""
}
