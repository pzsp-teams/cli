package parser

import (
	"strings"
	"testing"
)

func TestReadTemplate_ValidTemplate(t *testing.T) {
	template := "Hello {{.name}}! Welcome to {{.place}}."
	reader := strings.NewReader(template)

	tmpl, err := readTemplate(reader)
	if err != nil {
		t.Fatalf("ReadTemplate() unexpected error: %v", err)
	}

	if tmpl == nil {
		t.Error("ReadTemplate() returned nil template")
	}
}

func TestReadTemplate_InvalidSyntax(t *testing.T) {
	template := "Hello {{.name}! Missing closing braces"
	reader := strings.NewReader(template)

	_, err := readTemplate(reader)
	if err == nil {
		t.Error("ReadTemplate() expected error for invalid template syntax, got nil")
	}
}

func TestReadTemplate_UnclosedAction(t *testing.T) {
	template := "Hello {{.name"
	reader := strings.NewReader(template)

	_, err := readTemplate(reader)
	if err == nil {
		t.Error("ReadTemplate() expected error for unclosed action, got nil")
	}
}

func TestReadTemplate_InvalidAction(t *testing.T) {
	template := "Hello {{if .name}}"
	reader := strings.NewReader(template)

	_, err := readTemplate(reader)
	if err == nil {
		t.Error("ReadTemplate() expected error for unclosed if statement, got nil")
	}
}

func TestReadTemplate_EmptyTemplate(t *testing.T) {
	template := ""
	reader := strings.NewReader(template)

	tmpl, err := readTemplate(reader)
	if err != nil {
		t.Fatalf("ReadTemplate() unexpected error for empty template: %v", err)
	}

	if tmpl == nil {
		t.Error("ReadTemplate() returned nil template for empty string")
	}
}

func TestReadTemplate_MultilinePlaceholders(t *testing.T) {
	template := `Hello {{.name}}!

This is a multiline template.
Date: {{.date}}
Time: {{.time}}

Regards`
	reader := strings.NewReader(template)

	tmpl, err := readTemplate(reader)
	if err != nil {
		t.Fatalf("ReadTemplate() unexpected error: %v", err)
	}

	if tmpl == nil {
		t.Error("ReadTemplate() returned nil template")
	}
}
