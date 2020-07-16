package logger

import (
	"context"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"sync"
)

var once sync.Once
var instance *zap.Logger

func Configure() error {
	var err error
	once.Do(func() {
		instance, err = zap.NewProduction()
	})

	return err
}

func Debug(msg string, fields ...zap.Field) {
	instance.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	instance.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	instance.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	instance.Error(msg, fields...)
}

func ErrorWithContext(ctx context.Context, msg string, fields ...zap.Field) {
	ctxRqId, ok := ctx.Value(middleware.RequestIDKey).(string)
	if !ok {
		ctxRqId = "undefined"
	}

	zapFields := make([]zap.Field, 1+len(fields))
	zapFields[0] = zap.String("requestId", ctxRqId)
	zapFields = append(zapFields, fields...)
	instance.Error(msg, zapFields...)
}

func DPanic(msg string, fields ...zap.Field) {
	instance.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	instance.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	instance.Fatal(msg, fields...)
}
