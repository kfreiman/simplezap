package logging

import (
	"log"

	"go.uber.org/zap"
)

// Logger is a simplified interface over zap.Logger. https://github.com/uber-go/zap/issues/381
type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Fatal(args ...interface{})
}

// NewLogger creates zap's Logger implementation
func NewLogger(isDebug bool) Logger {
	config = zap.NewDevelopmentConfig()
	if !isDebug {
		atom := zap.NewAtomicLevel()
		atom.SetLevel(zap.InfoLevel)
		config = zap.NewProductionConfig()
	}

	logger, err := config.Build()
	if err != nil {
		log.Fatal(err)
	}

	return logger.Sugar()
}

type nopLogger struct {
}

// NewNopLogger creates Logger implementation which do nothing
func NewNopLogger() Logger {
	return &nopLogger{}
}

func (l *nopLogger) Debugw(msg string, keysAndValues ...interface{}) {
	// do nothing
}

func (l *nopLogger) Infow(msg string, keysAndValues ...interface{}) {
	// do nothing
}

func (l *nopLogger) Errorw(msg string, keysAndValues ...interface{}) {
	// do nothing
}

func (l *nopLogger) Fatal(args ...interface{}) {
	// do nothing
}
