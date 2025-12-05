package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestNewCharmFromConfig_DefaultConfig(t *testing.T) {
	log := NewCharmFromConfig(nil)
	if log == nil {
		t.Fatal("expected non-nil logger")
	}
}

func TestNewCharmFromConfig_CustomConfig(t *testing.T) {
	var buf bytes.Buffer
	cfg := &Config{
		Level:     LevelDebug,
		Format:    FormatText,
		Output:    &buf,
		AddSource: true,
	}

	log := NewCharmFromConfig(cfg)
	if log == nil {
		t.Fatal("expected non-nil logger")
	}

	log.Info("test message", "key", "value")

	output := buf.String()
	if !strings.Contains(output, "test message") {
		t.Errorf("expected output to contain 'test message', got: %s", output)
	}
	if !strings.Contains(output, "key") {
		t.Errorf("expected output to contain 'key', got: %s", output)
	}
}

func TestCharmLogger_AllLevels(t *testing.T) {
	var buf bytes.Buffer
	cfg := &Config{
		Level:  LevelDebug,
		Format: FormatText,
		Output: &buf,
	}

	log := NewCharmFromConfig(cfg)

	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")

	output := buf.String()
	if !strings.Contains(output, "debug message") {
		t.Errorf("expected debug message in output")
	}
	if !strings.Contains(output, "info message") {
		t.Errorf("expected info message in output")
	}
	if !strings.Contains(output, "warn message") {
		t.Errorf("expected warn message in output")
	}
	if !strings.Contains(output, "error message") {
		t.Errorf("expected error message in output")
	}
}

func TestCharmLogger_LevelFiltering(t *testing.T) {
	var buf bytes.Buffer
	cfg := &Config{
		Level:  LevelError,
		Format: FormatText,
		Output: &buf,
	}

	log := NewCharmFromConfig(cfg)

	log.Debug("debug message")
	log.Info("info message")
	log.Warn("warn message")
	log.Error("error message")

	output := buf.String()
	if strings.Contains(output, "debug message") {
		t.Errorf("debug message should not appear with Error level")
	}
	if strings.Contains(output, "info message") {
		t.Errorf("info message should not appear with Error level")
	}
	if strings.Contains(output, "warn message") {
		t.Errorf("warn message should not appear with Error level")
	}
	if !strings.Contains(output, "error message") {
		t.Errorf("error message should appear with Error level")
	}
}

func TestCharmLogger_With(t *testing.T) {
	var buf bytes.Buffer
	cfg := &Config{
		Level:  LevelInfo,
		Format: FormatText,
		Output: &buf,
	}

	log := NewCharmFromConfig(cfg)
	userLogger := log.With("user_id", "12345", "session", "abc-def")

	userLogger.Info("user action")

	output := buf.String()
	if !strings.Contains(output, "user_id") {
		t.Errorf("expected user_id in output, got: %s", output)
	}
	if !strings.Contains(output, "12345") {
		t.Errorf("expected 12345 in output, got: %s", output)
	}
	if !strings.Contains(output, "session") {
		t.Errorf("expected session in output, got: %s", output)
	}
	if !strings.Contains(output, "abc-def") {
		t.Errorf("expected abc-def in output, got: %s", output)
	}
}

func TestCharmLogger_WithChaining(t *testing.T) {
	var buf bytes.Buffer
	cfg := &Config{
		Level:  LevelInfo,
		Format: FormatText,
		Output: &buf,
	}

	log := NewCharmFromConfig(cfg)
	userLogger := log.With("user_id", "12345")
	requestLogger := userLogger.With("request_id", "req-789")

	requestLogger.Info("processing")

	output := buf.String()
	if !strings.Contains(output, "user_id") || !strings.Contains(output, "12345") {
		t.Errorf("expected user_id=12345 in output, got: %s", output)
	}
	if !strings.Contains(output, "request_id") || !strings.Contains(output, "req-789") {
		t.Errorf("expected request_id=req-789 in output, got: %s", output)
	}
}

func TestCharmLogger_ShowTimestamp(t *testing.T) {
	var bufWithTimestamp bytes.Buffer
	cfgWithTimestamp := &Config{
		Level:  LevelInfo,
		Format: FormatText,
		Output: &bufWithTimestamp,
	}

	logWithTimestamp := NewCharmFromConfig(cfgWithTimestamp)
	logWithTimestamp.Info("message with timestamp")

	outputWithTimestamp := bufWithTimestamp.String()

	var bufWithoutTimestamp bytes.Buffer
	cfgWithoutTimestamp := &Config{
		Level:         LevelInfo,
		Format:        FormatText,
		Output:        &bufWithoutTimestamp,
		OmitTimestamp: true,
	}

	logWithoutTimestamp := NewCharmFromConfig(cfgWithoutTimestamp)
	logWithoutTimestamp.Info("message without timestamp")

	outputWithoutTimestamp := bufWithoutTimestamp.String()

	if !strings.Contains(outputWithTimestamp, "message with timestamp") {
		t.Errorf("expected message in output with timestamp")
	}
	if !strings.Contains(outputWithoutTimestamp, "message without timestamp") {
		t.Errorf("expected message in output without timestamp")
	}

	if !strings.HasPrefix(strings.TrimSpace(outputWithoutTimestamp), "INFO") {
		t.Errorf("expected output without timestamp to start with INFO, got: %s", outputWithoutTimestamp)
	}
	if strings.HasPrefix(strings.TrimSpace(outputWithTimestamp), "INFO") {
		t.Errorf("expected output with timestamp to NOT start with INFO (should start with timestamp), got: %s", outputWithTimestamp)
	}
}
