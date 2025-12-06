package messages

import (
	"io"

	"github.com/pzsp-teams/cli/internal/initializers"
	"gopkg.in/yaml.v3"
)

// YAMLParser implements Parser for YAML format
type YAMLParser struct{}

// Parse reads YAML-formatted message data
func (p *YAMLParser) Parse(r io.Reader) (map[string]MessageData, error) {
	initializers.Logger.Debug("Parsing YAML data")
	var messages map[string]MessageData
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(&messages); err != nil {
		initializers.Logger.Error("Failed to decode YAML data", "error", err)
		return nil, err
	}
	initializers.Logger.Debug("Successfully parsed YAML data", "recipient_count", len(messages))
	return messages, nil
}
