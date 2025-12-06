package messages

import (
	"io"
	"text/template"

	"github.com/pzsp-teams/cli/internal/initializers"
)

// readTemplate reads template content from r and returns a parsed text/template.
// The template uses Go's text/template syntax with {{.placeholder}} format.
// Returns an error if reading fails or if the template syntax is invalid.
// Templates are configured to return an error if any placeholder is missing from the data.
func readTemplate(r io.Reader) (*template.Template, error) {
	initializers.Logger.Debug("Reading template content")
	content, err := io.ReadAll(r)
	if err != nil {
		initializers.Logger.Error("Failed to read template content", "error", err)
		return nil, err
	}
	initializers.Logger.Debug("Template content read", "size_bytes", len(content))

	tmpl, err := template.New("message").Option("missingkey=error").Parse(string(content))
	if err != nil {
		initializers.Logger.Error("Failed to parse template syntax", "error", err)
		return nil, err
	}

	initializers.Logger.Debug("Template parsed successfully")
	return tmpl, nil
}
