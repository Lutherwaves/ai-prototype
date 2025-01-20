// internal/platform/logger/logger.go
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Add caller and stacktrace for error levels
	config.EncoderConfig.StacktraceKey = "stacktrace"
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	logger, err := config.Build(
		zap.AddCallerSkip(1),
	)
	if err != nil {
		panic(err)
	}

	return logger
}
