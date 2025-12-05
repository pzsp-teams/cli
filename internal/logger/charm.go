package logger

import (
	"os"

	"github.com/charmbracelet/log"
)

type CharmLogger struct {
	logger *log.Logger
}

func NewCharmLogger(logger *log.Logger) *CharmLogger {
	if logger == nil {
		logger = log.NewWithOptions(os.Stderr, log.Options{})
	}
	return &CharmLogger{
		logger: logger,
	}
}

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

func (l *CharmLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *CharmLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *CharmLogger) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

func (l *CharmLogger) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

func (l *CharmLogger) With(args ...any) Logger {
	return &CharmLogger{
		logger: l.logger.With(args...),
	}
}
