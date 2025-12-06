package parser

import (
	"encoding/json"
	"io"
)

// JSONParser implements Parser for JSON format
type JSONParser struct{}

// Parse reads JSON-formatted message data
func (p *JSONParser) Parse(r io.Reader) (map[string]MessageData, error) {
	var messages map[string]MessageData
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&messages); err != nil {
		return nil, err
	}
	return messages, nil
}
