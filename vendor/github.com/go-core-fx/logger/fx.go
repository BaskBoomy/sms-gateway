package logger

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// WithFxDefaultLogger returns an fx.Option that configures Fx to use zap-based logging
// for framework events. The logger is set to DebugLevel to capture detailed
// Fx lifecycle and dependency injection events.
func WithFxDefaultLogger() fx.Option {
	return fx.WithLogger(func(logger *zap.Logger) fxevent.Logger {
		logOption := fxevent.ZapLogger{Logger: logger}
		logOption.UseLogLevel(zapcore.DebugLevel)
		return &logOption
	})
}

// WithNamedLogger returns an fx.Option that decorates the provided logger
// with the specified name, creating a named logger context.
func WithNamedLogger(name string) fx.Option {
	return fx.Decorate(func(logger *zap.Logger) *zap.Logger {
		return logger.Named(name)
	})
}
