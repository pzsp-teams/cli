package messages

import (
	"io"
	"text/template"
)

// readTemplate reads template content from r and returns a parsed text/template.
// The template uses Go's text/template syntax with {{.placeholder}} format.
// Returns an error if reading fails or if the template syntax is invalid.
// Templates are configured to return an error if any placeholder is missing from the data.
func readTemplate(r io.Reader) (*template.Template, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("message").Option("missingkey=error").Parse(string(content))
	if err != nil {
		return nil, err
	}

	return tmpl, nil
}
