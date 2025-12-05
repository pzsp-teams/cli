package main

import (
	"fmt"
	"os"

	"github.com/pzsp-teams/cli/internal/initializers"
	"github.com/pzsp-teams/cli/internal/logger"
)

func init() {
	verbose := false
	for _, arg := range os.Args {
		if arg == "-v" || arg == "--verbose" {
			verbose = true
			break
		}
	}

	logFile, err := os.Create("preview.log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create log file: %v\n", err)
		os.Exit(1)
	}

	stderrLevel := logger.LevelError
	if verbose {
		stderrLevel = logger.LevelDebug
	}

	initializers.InitMultiOutputLogger(initializers.MultiOutputConfig{
		StderrLevel:         stderrLevel,
		FileLevel:           logger.LevelDebug,
		FileWriter:          logFile,
		StderrOmitTimestamp: !verbose,
		FileOmitTimestamp:   false,
	})
}

func main() {
	charmDemo()
}

func charmDemo() {
	log := initializers.Logger
	log.Debug("this is a debug message")
	log.Info("application started successfully")
	log.Warn("this is a warning", "code", "WARN001")
	log.Error("this is an error", "error", "something went wrong")

	log.Info("=== Logging with Context (With) ===")
	userLogger := log.With("user_id", "12345", "role", "admin")
	userLogger.Info("user logged in")
	userLogger.Info("user accessed dashboard")

	requestLogger := userLogger.With("request_id", "req-abc-123")
	requestLogger.Info("processing request", "endpoint", "/api/users")
	requestLogger.Warn("slow request detected", "duration_ms", 2500)

	log.Debug("debug level message")
	log.Info("info level message")
	log.Warn("warning level message")
	log.Error("error level message")
}
