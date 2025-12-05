package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

// CharmLogger is a Logger implementation that wraps charmbracelet/log.Logger.
type CharmLogger struct {
	logger *log.Logger
}

// NewCharmLogger creates a new CharmLogger with the given charmbracelet log.Logger.
// If logger is nil, it uses a default logger that writes to stderr.
func NewCharmLogger(logger *log.Logger) *CharmLogger {
	if logger == nil {
		logger = log.NewWithOptions(os.Stderr, log.Options{})
	}
	return &CharmLogger{
		logger: logger,
	}
}

// NewCharmFromConfig creates a new CharmLogger with given configuration.
// This function handles the conversion from generic Config to charmbracelet-specific types.
func NewCharmFromConfig(cfg *Config) Logger {
	if cfg == nil {
		cfg = DefaultConfig()
	}

	if cfg.Output == nil {
		cfg.Output = os.Stderr
	}

	var charmLevel log.Level
	switch cfg.Level {
	case LevelDebug:
		charmLevel = log.DebugLevel
	case LevelInfo:
		charmLevel = log.InfoLevel
	case LevelWarn:
		charmLevel = log.WarnLevel
	case LevelError:
		charmLevel = log.ErrorLevel
	default:
		charmLevel = log.InfoLevel
	}

	opts := log.Options{
		Level:           charmLevel,
		ReportTimestamp: true,
		ReportCaller:    cfg.AddSource,
	}

	charmLogger := log.NewWithOptions(cfg.Output, opts)

	return NewCharmLogger(charmLogger)
}

// Debug logs a debug-level message with optional key-value pairs.
func (l *CharmLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Info logs an info-level message with optional key-value pairs.
func (l *CharmLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Warn logs a warning-level message with optional key-value pairs.
func (l *CharmLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Error logs an error-level message with optional key-value pairs.
func (l *CharmLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

// With returns a new Logger with the given key-value pairs added to the context.
func (l *CharmLogger) With(args ...any) Logger {
	return &CharmLogger{
		logger: l.logger.With(args...),
	}
}
