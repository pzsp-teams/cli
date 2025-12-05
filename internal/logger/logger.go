package logger

// Logger defines the interface for logging operations.
type Logger interface {
	// Debug logs a debug-level message with optional key-value pairs.
	Debug(msg string, args ...any)

	// Info logs an info-level message with optional key-value pairs.
	Info(msg string, args ...any)

	// Warn logs a warning-level message with optional key-value pairs.
	Warn(msg string, args ...any)

	// Error logs an error-level message with optional key-value pairs.
	Error(msg string, args ...any)

	// With returns a new Logger with the given key-value pairs added to the context.
	With(args ...any) Logger
}
