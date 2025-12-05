package logger

import (
	"io"
	"os"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
)

type Format int

const (
	FormatText Format = iota
	FormatJSON
)

type Config struct {
	Level     Level
	Format    Format
	Output    io.Writer
	AddSource bool
}

func DefaultConfig() *Config {
	return &Config{
		Level:     LevelInfo,
		Format:    FormatText,
		Output:    os.Stdout,
		AddSource: false,
	}
}
