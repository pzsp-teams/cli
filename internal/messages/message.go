package messages

import (
	"bytes"
	"fmt"
	"io"

	"github.com/pzsp-teams/cli/internal/initializers"
)

// MessageParser handles parsing different messages from supplied template and data
type MessageParser struct {
	templateReader io.Reader
	dataReader     io.Reader
	dataParser     Parser
}

// NewMessageParser returns a MessageParser with given config
func NewMessageParser(templateReader, dataReader io.Reader, dataParser Parser) *MessageParser {
	return &MessageParser{
		templateReader: templateReader,
		dataReader:     dataReader,
		dataParser:     dataParser,
	}
}

// Parse reads the template and data files, then returns a map of rendered messages.
// The map keys are recipient names, and values are the fully rendered messages.
func (mp *MessageParser) Parse() (map[string]string, error) {
	initializers.Logger.Debug("Starting message parsing")

	tmpl, err := readTemplate(mp.templateReader)
	if err != nil {
		initializers.Logger.Error("Failed to read template", "error", err)
		return nil, fmt.Errorf("failed to read template: %w", err)
	}
	initializers.Logger.Debug("Template read successfully")

	recipients, err := mp.dataParser.Parse(mp.dataReader)
	if err != nil {
		initializers.Logger.Error("Failed to parse message data", "error", err)
		return nil, fmt.Errorf("failed to parse message data: %w", err)
	}
	initializers.Logger.Info("Parsed message data", "recipient_count", len(recipients))

	messages := make(map[string]string, len(recipients))
	for recipientName, data := range recipients {
		initializers.Logger.Debug("Rendering message for recipient", "recipient", recipientName)
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			initializers.Logger.Error("Failed to render message", "recipient", recipientName, "error", err)
			return nil, fmt.Errorf("failed to render message for recipient %q: %w", recipientName, err)
		}
		messages[recipientName] = buf.String()
	}

	initializers.Logger.Info("Successfully generated messages", "total_messages", len(messages))
	return messages, nil
}
