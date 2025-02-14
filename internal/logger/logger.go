// Package logger provides logging functionality for the application
package logger

import (
	"context"

	"github.com/LabGroupware/go-measure-tui/internal/config"
)

// Logger is an interface for logging
type Logger interface {
	SetupLogger(config *config.LoggingConfig) error
	With(args ...keyVal) Logger
	Debug(ctx context.Context, msg string, args ...keyVal)
	Info(ctx context.Context, msg string, args ...keyVal)
	Warn(ctx context.Context, msg string, args ...keyVal)
	Error(ctx context.Context, msg string, args ...keyVal)
	Close() error
}

type keyVal struct {
	Key   string
	Value any
}

// Value creates a new KeyVal
func Value(key string, value any) keyVal {
	return keyVal{Key: key, Value: value}
}

// Group creates a new KeyVal with a group of KeyVals
func Group(key string, kvs ...keyVal) keyVal {
	return keyVal{Key: key, Value: kvs}
}
