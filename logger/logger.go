package logger

import (
	"log"

	"go.uber.org/zap"
)

type Logger interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	With(args ...interface{}) *zap.SugaredLogger
}

func NewLogger() Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("error initialising logger: %+v", err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	return logger.Sugar()
}

func NewNopLogger() Logger {
	return zap.NewNop().Sugar()
}
