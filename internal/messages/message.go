package messages

import (
	"bytes"
	"fmt"
	"io"
	"text/template"

	"github.com/pzsp-teams/cli/internal/initializers"
)

// MessageParser handles parsing different messages from supplied template and data
type MessageParser struct {
	template   *template.Template
	recipients map[string]MessageData
}

// NewMessageParser returns a MessageParser with given config.
// It parses the template and data immediately, storing the parsed objects.
func NewMessageParser(templateReader, dataReader io.Reader, dataParser Parser) (*MessageParser, error) {
	initializers.Logger.Debug("Creating new MessageParser")

	tmpl, err := readTemplate(templateReader)
	if err != nil {
		initializers.Logger.Error("Failed to read and parse template", "error", err)
		return nil, fmt.Errorf("failed to read and parse template: %w", err)
	}
	initializers.Logger.Debug("Template parsed successfully")

	recipients, err := dataParser.Parse(dataReader)
	if err != nil {
		initializers.Logger.Error("Failed to parse message data", "error", err)
		return nil, fmt.Errorf("failed to parse message data: %w", err)
	}
	initializers.Logger.Info("Message data parsed", "recipient_count", len(recipients))

	return &MessageParser{
		template:   tmpl,
		recipients: recipients,
	}, nil
}

// Parse renders the template for each recipient and returns a map of rendered messages.
// The map keys are recipient names, and values are the fully rendered messages.
func (mp *MessageParser) Parse() (map[string]string, error) {
	initializers.Logger.Debug("Starting message rendering")

	messages := make(map[string]string, len(mp.recipients))
	for recipientName, data := range mp.recipients {
		initializers.Logger.Debug("Rendering message for recipient", "recipient", recipientName)
		var buf bytes.Buffer
		if err := mp.template.Execute(&buf, data); err != nil {
			initializers.Logger.Error("Failed to render message", "recipient", recipientName, "error", err)
			return nil, fmt.Errorf("failed to render message for recipient %q: %w", recipientName, err)
		}
		messages[recipientName] = buf.String()
	}

	initializers.Logger.Info("Successfully rendered messages", "total_messages", len(messages))
	return messages, nil
}
