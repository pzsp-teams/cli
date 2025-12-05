package logger

import (
	"bytes"
	"os"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg == nil {
		t.Fatal("expected non-nil config")
	}

	if cfg.Level != LevelInfo {
		t.Errorf("expected default level to be LevelInfo, got %v", cfg.Level)
	}

	if cfg.Format != FormatText {
		t.Errorf("expected default format to be FormatText, got %v", cfg.Format)
	}

	if cfg.Output != os.Stdout {
		t.Errorf("expected default output to be os.Stdout")
	}

	if cfg.AddSource {
		t.Errorf("expected default AddSource to be false")
	}

	if cfg.OmitTimestamp {
		t.Errorf("expected default OmitTimestamp to be false")
	}
}

func TestConfig_CustomValues(t *testing.T) {
	var buf bytes.Buffer
	cfg := &Config{
		Level:         LevelDebug,
		Format:        FormatJSON,
		Output:        &buf,
		AddSource:     true,
		OmitTimestamp: true,
	}

	if cfg.Level != LevelDebug {
		t.Errorf("expected level to be LevelDebug, got %v", cfg.Level)
	}

	if cfg.Format != FormatJSON {
		t.Errorf("expected format to be FormatJSON, got %v", cfg.Format)
	}

	if cfg.Output != &buf {
		t.Errorf("expected output to be custom buffer")
	}

	if !cfg.AddSource {
		t.Errorf("expected AddSource to be true")
	}

	if !cfg.OmitTimestamp {
		t.Errorf("expected OmitTimestamp to be true")
	}
}
