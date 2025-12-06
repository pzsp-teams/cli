package parser

import (
	"bytes"
	"fmt"
	"io"
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
	tmpl, err := readTemplate(mp.templateReader)
	if err != nil {
		return nil, fmt.Errorf("failed to read template: %w", err)
	}

	recipients, err := mp.dataParser.Parse(mp.dataReader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse message data: %w", err)
	}

	messages := make(map[string]string, len(recipients))
	for recipientName, data := range recipients {
		var buf bytes.Buffer
		if err := tmpl.Execute(&buf, data); err != nil {
			return nil, fmt.Errorf("failed to render message for recipient %q: %w", recipientName, err)
		}
		messages[recipientName] = buf.String()
	}

	return messages, nil
}
