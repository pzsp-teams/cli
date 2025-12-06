package messages

import (
	"io"

	"github.com/BurntSushi/toml"
	"github.com/pzsp-teams/cli/internal/initializers"
)

// TOMLParser implements Parser for TOML format
type TOMLParser struct{}

// Parse reads TOML-formatted message data
func (p *TOMLParser) Parse(r io.Reader) (map[string]MessageData, error) {
	var messages map[string]MessageData
	if _, err := toml.NewDecoder(r).Decode(&messages); err != nil {
		initializers.Logger.Error("Failed to decode TOML data", "error", err)
		return nil, err
	}
	return messages, nil
}
