package logger

import (
	"io"
	"os"
)

// Level represents the logging level.
type Level int

// Available logging levels
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

// Format represents the log output format.
type Format int

// Available log output formats.
const (
	FormatText Format = iota
	FormatJSON
)

// Config holds the configuration for creating a logger.
type Config struct {
	// Level sets the minimum log level. Messages below this level are discarded.
	Level Level
	// Format sets the output format (text or JSON).
	Format Format
	// Output sets where logs are written. If nil, defaults to os.Stderr.
	Output io.Writer
	// AddSource adds source code position (file and line number) to log records.
	AddSource bool
}

// DefaultConfig returns a Config with default settings:
//
//	Level:     LevelInfo
//	Format:    FormatText
//	Output:    os.Stdout
//	AddSource: false
func DefaultConfig() *Config {
	return &Config{
		Level:     LevelInfo,
		Format:    FormatText,
		Output:    os.Stdout,
		AddSource: false,
	}
}
