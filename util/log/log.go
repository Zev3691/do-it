package log

import (
	"context"
	"re_new/util"

	"go.uber.org/zap"
)

var logg *zap.Logger

func Init() {
	switch util.GetVersion() {
	case util.Production:
		logg, _ = zap.NewProduction()
		logg.Info(util.Production)
	case util.Development:
		logg, _ = zap.NewDevelopment()
		logg.Info(util.Development)
	default:
		logg = zap.NewExample()
		logg.Info(util.Test)
	}
}

func Info(msg string, fields ...zap.Field) {
	logg.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logg.Debug(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logg.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logg.Fatal(msg, fields...)
}

func Painc(msg string, fields ...zap.Field) {
	logg.Panic(msg, fields...)
}

// 自定义格式化输出
func Infof(msg string, data ...interface{}) {
	logg.Sugar().Infof(msg, data...)
}

func Debugf(msg string, data ...interface{}) {
	logg.Sugar().Debugf(msg, data...)
}

func Errorf(msg string, data ...interface{}) {
	logg.Sugar().Errorf(msg, data...)
}

func Fatalf(msg string, data ...interface{}) {
	logg.Sugar().Fatalf(msg, data...)
}

func Paincf(msg string, data ...interface{}) {
	logg.Sugar().Panicf(msg, data...)
}

func FromContext(ctx context.Context) *zap.Logger {
	reqId := ctx.Value("request_id").(string)
	return logg.With(zap.String("type", "gorm"), zap.String("request_id", reqId))
}
