package log

import (
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
