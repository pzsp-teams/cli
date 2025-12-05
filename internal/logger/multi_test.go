package logger

import (
	"bytes"
	"strings"
	"testing"
)

func TestMultiLogger_DispatchesToAll(t *testing.T) {
	var buf1, buf2 bytes.Buffer

	logger1 := NewCharmFromConfig(&Config{
		Level:  LevelDebug,
		Format: FormatText,
		Output: &buf1,
	})

	logger2 := NewCharmFromConfig(&Config{
		Level:  LevelDebug,
		Format: FormatText,
		Output: &buf2,
	})

	multi := NewMultiLogger(logger1, logger2)
	multi.Info("test message")

	output1 := buf1.String()
	output2 := buf2.String()

	if !strings.Contains(output1, "test message") {
		t.Errorf("logger1 did not receive message, got: %s", output1)
	}

	if !strings.Contains(output2, "test message") {
		t.Errorf("logger2 did not receive message, got: %s", output2)
	}
}

func TestMultiLogger_RespectsIndividualLevels(t *testing.T) {
	var bufDebug, bufError bytes.Buffer

	debugLogger := NewCharmFromConfig(&Config{
		Level:  LevelDebug,
		Format: FormatText,
		Output: &bufDebug,
	})

	errorLogger := NewCharmFromConfig(&Config{
		Level:  LevelError,
		Format: FormatText,
		Output: &bufError,
	})

	multi := NewMultiLogger(debugLogger, errorLogger)

	multi.Debug("debug message")
	multi.Info("info message")
	multi.Error("error message")

	debugOutput := bufDebug.String()
	errorOutput := bufError.String()

	if !strings.Contains(debugOutput, "debug message") {
		t.Errorf("debug logger should have debug message")
	}
	if !strings.Contains(debugOutput, "info message") {
		t.Errorf("debug logger should have info message")
	}
	if !strings.Contains(debugOutput, "error message") {
		t.Errorf("debug logger should have error message")
	}

	if strings.Contains(errorOutput, "debug message") {
		t.Errorf("error logger should not have debug message")
	}
	if strings.Contains(errorOutput, "info message") {
		t.Errorf("error logger should not have info message")
	}
	if !strings.Contains(errorOutput, "error message") {
		t.Errorf("error logger should have error message")
	}
}

func TestMultiLogger_With(t *testing.T) {
	var buf1, buf2 bytes.Buffer

	logger1 := NewCharmFromConfig(&Config{
		Level:  LevelInfo,
		Format: FormatText,
		Output: &buf1,
	})

	logger2 := NewCharmFromConfig(&Config{
		Level:  LevelInfo,
		Format: FormatText,
		Output: &buf2,
	})

	multi := NewMultiLogger(logger1, logger2)
	contextLogger := multi.With("user_id", "123")
	contextLogger.Info("test message")

	output1 := buf1.String()
	output2 := buf2.String()

	if !strings.Contains(output1, "user_id") || !strings.Contains(output1, "123") {
		t.Errorf("logger1 should have context, got: %s", output1)
	}

	if !strings.Contains(output2, "user_id") || !strings.Contains(output2, "123") {
		t.Errorf("logger2 should have context, got: %s", output2)
	}
}

func TestMultiLogger_NoLoggers(t *testing.T) {
	multi := NewMultiLogger()

	multi.Debug("test")
	multi.Info("test")
	multi.Warn("test")
	multi.Error("test")
	multi.With("key", "value").Info("test")
}

func TestMultiLogger_SingleLogger(t *testing.T) {
	var buf bytes.Buffer

	logger := NewCharmFromConfig(&Config{
		Level:  LevelInfo,
		Format: FormatText,
		Output: &buf,
	})

	multi := NewMultiLogger(logger)
	multi.Info("test message")

	output := buf.String()
	if !strings.Contains(output, "test message") {
		t.Errorf("expected message in output, got: %s", output)
	}
}
