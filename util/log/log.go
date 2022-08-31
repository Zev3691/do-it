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

func GetOriginLog() *zap.Logger {
	return logg
}

func Info(ctx context.Context, msg string, fields ...zap.Field) {
	log := FromContext(ctx)
	log.Info(msg, fields...)
}

func Debug(ctx context.Context, msg string, fields ...zap.Field) {
	log := FromContext(ctx)
	log.Debug(msg, fields...)
}

func Error(ctx context.Context, msg string, fields ...zap.Field) {
	log := FromContext(ctx)
	log.Error(msg, fields...)
}

func Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	log := FromContext(ctx)
	log.Fatal(msg, fields...)
}

func Painc(ctx context.Context, msg string, fields ...zap.Field) {
	log := FromContext(ctx)
	log.Panic(msg, fields...)
}

// 自定义格式化输出
func Infof(ctx context.Context, msg string, data ...interface{}) {
	log := FromContext(ctx)
	log.Sugar().Infof(msg, data...)
}

func Debugf(ctx context.Context, msg string, data ...interface{}) {
	log := FromContext(ctx)
	log.Sugar().Debugf(msg, data...)
}

func Errorf(ctx context.Context, msg string, data ...interface{}) {
	log := FromContext(ctx)
	log.Sugar().Errorf(msg, data...)
}

func Fatalf(ctx context.Context, msg string, data ...interface{}) {
	log := FromContext(ctx)
	log.Sugar().Fatalf(msg, data...)
}

func Paincf(ctx context.Context, msg string, data ...interface{}) {
	log := FromContext(ctx)
	log.Sugar().Panicf(msg, data...)
}

func FromContext(ctx context.Context) *zap.Logger {
	reqId := ctx.Value("request_id")
	if reqId == nil {
		return logg
	}
	return logg.With(zap.String("request_id", reqId.(string)))
}
