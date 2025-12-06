package messages

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pzsp-teams/cli/internal/initializers"
)

// Registry manages available parsers for different file formats
type Registry struct {
	parsers map[string]Parser
}

// NewParserRegistry creates a new registry with default parsers
func NewParserRegistry() *Registry {
	initializers.Logger.Debug("Creating new parser registry")
	registry := &Registry{
		parsers: make(map[string]Parser),
	}

	registry.Register("json", &JSONParser{})
	registry.Register("yaml", &YAMLParser{})
	registry.Register("yml", &YAMLParser{})
	registry.Register("toml", &TOMLParser{})

	initializers.Logger.Info("Parser registry initialized", "supported_formats", registry.SupportedFormats())
	return registry
}

// Register adds a parser for a specific file extension
func (r *Registry) Register(ext string, parser Parser) {
	normalizedExt := strings.ToLower(ext)
	initializers.Logger.Debug("Registering parser", "extension", normalizedExt)
	r.parsers[normalizedExt] = parser
}

// GetParser returns a parser for the given file extension
func (r *Registry) GetParser(filename string) (Parser, error) {
	ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")
	initializers.Logger.Debug("Looking up parser for file", "filename", filename, "extension", ext)

	parser, ok := r.parsers[ext]
	if !ok {
		initializers.Logger.Warn("No parser found for extension", "extension", ext, "supported_formats", r.SupportedFormats())
		return nil, fmt.Errorf("no parser registered for extension: .%s", ext)
	}

	initializers.Logger.Debug("Parser found", "extension", ext)
	return parser, nil
}

// SupportedFormats returns a list of supported file extensions
func (r *Registry) SupportedFormats() []string {
	formats := make([]string, 0, len(r.parsers))
	for ext := range r.parsers {
		formats = append(formats, ext)
	}
	return formats
}
