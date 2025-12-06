package templates

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"text/template"

	"github.com/pzsp-teams/cli/internal/initializers"
)

var htmlTagRegex = regexp.MustCompile(`<\/?(?:i|b|p)>|<br>|<a\s+href="[^"]*">|<\/a>`)

// TemplateParser handles parsing different messages from supplied template and data
type TemplateParser struct {
	template   *template.Template
	recipients map[string]TemplateData
}

// NewMessageParser returns a MessageParser with given config.
// It parses the template and data immediately, storing the parsed objects.
func NewMessageParser(templateReader, dataReader io.Reader, dataParser Parser) (*TemplateParser, error) {
	tmpl, err := readTemplate(templateReader)
	if err != nil {
		initializers.Logger.Error("Failed to read and parse template", "error", err)
		return nil, fmt.Errorf("failed to read and parse template: %w", err)
	}

	recipients, err := dataParser.Parse(dataReader)
	if err != nil {
		initializers.Logger.Error("Failed to parse message data", "error", err)
		return nil, fmt.Errorf("failed to parse message data: %w", err)
	}
	initializers.Logger.Info("Message data parsed", "recipient_count", len(recipients))

	return &TemplateParser{
		template:   tmpl,
		recipients: recipients,
	}, nil
}

// Parse renders the template for each recipient and returns a map of rendered messages.
// The map keys are recipient names, and values are the fully rendered messages.
func (mp *TemplateParser) Parse() (map[string]string, error) {
	messages := make(map[string]string, len(mp.recipients))
	for recipientName, data := range mp.recipients {
		var buf bytes.Buffer
		if err := mp.template.Execute(&buf, data); err != nil {
			initializers.Logger.Error("Failed to render message", "recipient", recipientName, "error", err)
			return nil, fmt.Errorf("failed to render message for recipient %q: %w", recipientName, err)
		}
		messages[recipientName] = processContent(buf.Bytes())
	}

	initializers.Logger.Info("Successfully rendered messages", "total_messages", len(messages))
	return messages, nil
}

func processContent(data []byte) string {
	if htmlTagRegex.Match(data) {
		return string(data)
	}

	// \r can be ignored, won't be rendered by html on teams anyway
	replaced := bytes.ReplaceAll(data, []byte("\n"), []byte("<br>"))
	return string(replaced)
}
