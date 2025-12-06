package templates

import (
	"io"

	"github.com/pzsp-teams/cli/internal/initializers"
	"gopkg.in/yaml.v3"
)

// YAMLParser implements Parser for YAML format
type YAMLParser struct{}

// Parse reads YAML-formatted message data
func (p *YAMLParser) Parse(r io.Reader) (map[string]TemplateData, error) {
	var messages map[string]TemplateData
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(&messages); err != nil {
		initializers.Logger.Error("Failed to decode YAML data", "error", err)
		return nil, err
	}
	return messages, nil
}
