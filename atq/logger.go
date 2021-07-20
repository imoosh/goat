package atq

import (
	"github.com/hibiken/asynq"
	"goat/log"
)

const defaultLoggerLevel = asynq.InfoLevel

type logger struct {
}

// Debug logs a message at Debug level.
func (l logger) Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info logs a message at Info level.
func (l logger) Info(args ...interface{}) {
	log.Info(args...)
}

// Warn logs a message at Warning level.
func (l logger) Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error logs a message at Error level.
func (l logger) Error(args ...interface{}) {
	log.Error(args...)
}

// Fatal logs a message at Fatal level
// and process will exit with status set to 1.
func (l logger) Fatal(args ...interface{}) {
	log.Debug(args...)
}
