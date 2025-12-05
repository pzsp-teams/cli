package main

import (
	"fmt"
	"os"

	"github.com/pzsp-teams/cli/internal/logger"
)

func main() {
	charmDemo()
}

func charmDemo() {
	log := logger.NewCharmFromConfig(logger.DefaultConfig())

	log.Info("=== Basic Logging ===")
	log.Debug("this is a debug message (won't show with default config)")
	log.Info("application started successfully")
	log.Warn("this is a warning", "code", "WARN001")
	log.Error("this is an error", "error", "something went wrong")

	fmt.Println("")
	log.Info("=== Logging with Context (With) ===")
	userLogger := log.With("user_id", "12345", "role", "admin")
	userLogger.Info("user logged in")
	userLogger.Info("user accessed dashboard")

	requestLogger := userLogger.With("request_id", "req-abc-123")
	requestLogger.Info("processing request", "endpoint", "/api/users")
	requestLogger.Warn("slow request detected", "duration_ms", 2500)

	fmt.Println("")
	log.Info("=== Debug Level ===")
	debugLog := logger.NewCharmFromConfig(&logger.Config{
		Level:  logger.LevelDebug,
		Format: logger.FormatText,
		Output: os.Stdout,
	})
	debugLog.Debug("now you can see debug messages")
	debugLog.Info("info level message")
	debugLog.Warn("warning level message")
	debugLog.Error("error level message")

	fmt.Println("")
	noTimestampLog := logger.NewCharmFromConfig(&logger.Config{
		Level:         logger.LevelInfo,
		Format:        logger.FormatText,
		Output:        os.Stdout,
		OmitTimestamp: true,
	})
	noTimestampLog.Info("=== Logs Without Timestamps ===")
	noTimestampLog.Info("clean log without timestamp")
	noTimestampLog.Warn("warning without timestamp", "reason", "demonstration")
	noTimestampLog.Error("error without timestamp", "code", "ERR001")
}
