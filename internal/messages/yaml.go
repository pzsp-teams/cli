package messages

import (
	"io"

	"gopkg.in/yaml.v3"
)

// YAMLParser implements Parser for YAML format
type YAMLParser struct{}

// Parse reads YAML-formatted message data
func (p *YAMLParser) Parse(r io.Reader) (map[string]MessageData, error) {
	var messages map[string]MessageData
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(&messages); err != nil {
		return nil, err
	}
	return messages, nil
}
