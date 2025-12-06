package templates

import (
	"strings"
	"testing"
)

func TestMessageParser_JSONFormatMultipleRecipients(t *testing.T) {
	template := "Hello {{.name}}! Your order #{{.order}} is ready."
	data := `{
		"alice": {
			"name": "Alice",
			"order": "12345"
		},
		"bob": {
			"name": "Bob",
			"order": "67890"
		}
	}`

	tmplReader := strings.NewReader(template)
	dataReader := strings.NewReader(data)

	mp, err := NewMessageParser(tmplReader, dataReader, &JSONParser{})
	if err != nil {
		t.Fatalf("NewMessageParser() unexpected error: %v", err)
	}
	messages, err := mp.Parse()
	if err != nil {
		t.Fatalf("MessageParser.Parse() unexpected error: %v", err)
	}

	wantMessages := map[string]string{
		"alice": "Hello Alice! Your order #12345 is ready.",
		"bob":   "Hello Bob! Your order #67890 is ready.",
	}

	if len(messages) != len(wantMessages) {
		t.Errorf("MessageParser.Parse() got %d messages, want %d", len(messages), len(wantMessages))
	}

	for recipient, wantMsg := range wantMessages {
		gotMsg, ok := messages[recipient]
		if !ok {
			t.Errorf("MessageParser.Parse() missing message for recipient %q", recipient)
			continue
		}
		if gotMsg != wantMsg {
			t.Errorf("MessageParser.Parse() for recipient %q:\ngot:\n%s\nwant:\n%s", recipient, gotMsg, wantMsg)
		}
	}
}

func TestMessageParser_TemplateWithMultiplePlaceholders(t *testing.T) {
	template := `Hello {{.name}}!

Meeting reminder:
Date: {{.date}}
Time: {{.time}}
Topic: {{.topic}}

See you there!`
	data := `{
		"team1": {
			"name": "Team Alpha",
			"date": "December 1, 2025",
			"time": "10:00",
			"topic": "Q4 Review"
		},
		"team2": {
			"name": "Team Beta",
			"date": "December 1, 2025",
			"time": "11:00",
			"topic": "Project Review"
		}
	}`

	tmplReader := strings.NewReader(template)
	dataReader := strings.NewReader(data)

	mp, err := NewMessageParser(tmplReader, dataReader, &JSONParser{})
	if err != nil {
		t.Fatalf("NewMessageParser() unexpected error: %v", err)
	}
	messages, err := mp.Parse()
	if err != nil {
		t.Fatalf("MessageParser.Parse() unexpected error: %v", err)
	}

	wantMessages := map[string]string{
		"team1": `Hello Team Alpha!

Meeting reminder:
Date: December 1, 2025
Time: 10:00
Topic: Q4 Review

See you there!`,
		"team2": `Hello Team Beta!

Meeting reminder:
Date: December 1, 2025
Time: 11:00
Topic: Project Review

See you there!`,
	}

	if len(messages) != len(wantMessages) {
		t.Errorf("MessageParser.Parse() got %d messages, want %d", len(messages), len(wantMessages))
	}

	for recipient, wantMsg := range wantMessages {
		gotMsg, ok := messages[recipient]
		if !ok {
			t.Errorf("MessageParser.Parse() missing message for recipient %q", recipient)
			continue
		}
		if gotMsg != wantMsg {
			t.Errorf("MessageParser.Parse() for recipient %q:\ngot:\n%s\nwant:\n%s", recipient, gotMsg, wantMsg)
		}
	}
}

func TestMessageParser_YAMLFormat(t *testing.T) {
	template := "Hello {{.name}}!"
	data := `
alice:
  name: Alice
bob:
  name: Bob
`

	tmplReader := strings.NewReader(template)
	dataReader := strings.NewReader(data)

	mp, err := NewMessageParser(tmplReader, dataReader, &YAMLParser{})
	if err != nil {
		t.Fatalf("NewMessageParser() unexpected error: %v", err)
	}
	messages, err := mp.Parse()
	if err != nil {
		t.Fatalf("MessageParser.Parse() unexpected error: %v", err)
	}

	wantMessages := map[string]string{
		"alice": "Hello Alice!",
		"bob":   "Hello Bob!",
	}

	if len(messages) != len(wantMessages) {
		t.Errorf("MessageParser.Parse() got %d messages, want %d", len(messages), len(wantMessages))
	}

	for recipient, wantMsg := range wantMessages {
		gotMsg, ok := messages[recipient]
		if !ok {
			t.Errorf("MessageParser.Parse() missing message for recipient %q", recipient)
			continue
		}
		if gotMsg != wantMsg {
			t.Errorf("MessageParser.Parse() for recipient %q got %q, want %q", recipient, gotMsg, wantMsg)
		}
	}
}

func TestMessageParser_TOMLFormat(t *testing.T) {
	template := "Hello {{.name}}!"
	data := `
[alice]
name = "Alice"

[bob]
name = "Bob"
`

	tmplReader := strings.NewReader(template)
	dataReader := strings.NewReader(data)

	mp, err := NewMessageParser(tmplReader, dataReader, &TOMLParser{})
	if err != nil {
		t.Fatalf("NewMessageParser() unexpected error: %v", err)
	}
	messages, err := mp.Parse()
	if err != nil {
		t.Fatalf("MessageParser.Parse() unexpected error: %v", err)
	}

	wantMessages := map[string]string{
		"alice": "Hello Alice!",
		"bob":   "Hello Bob!",
	}

	if len(messages) != len(wantMessages) {
		t.Errorf("MessageParser.Parse() got %d messages, want %d", len(messages), len(wantMessages))
	}

	for recipient, wantMsg := range wantMessages {
		gotMsg, ok := messages[recipient]
		if !ok {
			t.Errorf("MessageParser.Parse() missing message for recipient %q", recipient)
			continue
		}
		if gotMsg != wantMsg {
			t.Errorf("MessageParser.Parse() for recipient %q got %q, want %q", recipient, gotMsg, wantMsg)
		}
	}
}

func TestMessageParser_InvalidTemplateSyntax(t *testing.T) {
	template := "Hello {{.name}"
	data := `{"alice": {"name": "Alice"}}`

	tmplReader := strings.NewReader(template)
	dataReader := strings.NewReader(data)

	_, err := NewMessageParser(tmplReader, dataReader, &JSONParser{})

	if err == nil {
		t.Error("NewMessageParser() expected error for invalid template syntax, got nil")
	}
}

func TestMessageParser_InvalidJSONData(t *testing.T) {
	template := "Hello {{.name}}!"
	data := `{"alice": invalid json}`

	tmplReader := strings.NewReader(template)
	dataReader := strings.NewReader(data)

	_, err := NewMessageParser(tmplReader, dataReader, &JSONParser{})

	if err == nil {
		t.Error("NewMessageParser() expected error for invalid JSON data, got nil")
	}
}

func TestMessageParser_InvalidYAMLData(t *testing.T) {
	template := "Hello {{.name}}!"
	data := `alice:
  name: Alice
    invalid: indentation
  email: test
`

	tmplReader := strings.NewReader(template)
	dataReader := strings.NewReader(data)

	_, err := NewMessageParser(tmplReader, dataReader, &YAMLParser{})

	if err == nil {
		t.Error("NewMessageParser() expected error for invalid YAML data, got nil")
	}
}

func TestMessageParser_InvalidTOMLData(t *testing.T) {
	template := "Hello {{.name}}!"
	data := `[alice]
name = "Alice"
email = invalid toml without quotes
`

	tmplReader := strings.NewReader(template)
	dataReader := strings.NewReader(data)

	_, err := NewMessageParser(tmplReader, dataReader, &TOMLParser{})

	if err == nil {
		t.Error("NewMessageParser() expected error for invalid TOML data, got nil")
	}
}

func TestMessageParser_MissingPlaceholderInData(t *testing.T) {
	template := "Hello {{.name}}! Your email is {{.email}}"
	data := `{"alice": {"name": "Alice"}}`

	tmplReader := strings.NewReader(template)
	dataReader := strings.NewReader(data)

	mp, err := NewMessageParser(tmplReader, dataReader, &JSONParser{})
	if err != nil {
		t.Fatalf("NewMessageParser() unexpected error: %v", err)
	}
	_, err = mp.Parse()

	if err == nil {
		t.Error("MessageParser.Parse() expected error for missing placeholder, got nil")
	}
}

func TestMessageParser_EmptyData(t *testing.T) {
	tmplReader := strings.NewReader("Hello {{.name}}!")
	dataReader := strings.NewReader(`{}`)

	mp, err := NewMessageParser(tmplReader, dataReader, &JSONParser{})
	if err != nil {
		t.Fatalf("NewMessageParser() unexpected error: %v", err)
	}
	messages, err := mp.Parse()
	if err != nil {
		t.Errorf("MessageParser.Parse() unexpected error: %v", err)
	}

	if len(messages) != 0 {
		t.Errorf("MessageParser.Parse() got %d messages, want 0", len(messages))
	}
}
