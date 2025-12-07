package templates

import "errors"

var (
	// Template errors
	ErrTemplateReadFailed   = errors.New("failed to read template content")
	ErrTemplateParseFailed  = errors.New("failed to parse template syntax")
	ErrTemplateRenderFailed = errors.New("failed to render template")

	// Data parsing errors
	ErrDataParseFailed  = errors.New("failed to parse message data")
	ErrJSONDecodeFailed = errors.New("failed to decode JSON data")
	ErrYAMLDecodeFailed = errors.New("failed to decode YAML data")
	ErrTOMLDecodeFailed = errors.New("failed to decode TOML data")

	// Registry errors
	ErrUnsupportedFormat  = errors.New("unsupported file format")
	ErrNoParserRegistered = errors.New("no parser registered for extension")
)
