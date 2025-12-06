package messages

import (
	"encoding/json"
	"io"

	"github.com/pzsp-teams/cli/internal/initializers"
)

// JSONParser implements Parser for JSON format
type JSONParser struct{}

// Parse reads JSON-formatted message data
func (p *JSONParser) Parse(r io.Reader) (map[string]MessageData, error) {
	var messages map[string]MessageData
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&messages); err != nil {
		initializers.Logger.Error("Failed to decode JSON data", "error", err)
		return nil, err
	}
	return messages, nil
}
