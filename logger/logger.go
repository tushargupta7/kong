package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger struct to hold log configuration
type Logger struct {
	logger *log.Logger
}

// NewLogger creates and returns a new logger instance
func NewLogger() *Logger {
	// Create log file or use stdout
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		logFile = os.Stdout
	}

	return &Logger{
		logger: log.New(logFile, "", log.LstdFlags),
	}
}

// Info logs an info level message
func (l *Logger) Info(message string, context map[string]interface{}) {
	l.log("INFO", message, context)
}

// Warn logs a warning level message
func (l *Logger) Warn(message string, context map[string]interface{}) {
	l.log("WARN", message, context)
}

// Error logs an error level message
func (l *Logger) Error(message string, context map[string]interface{}) {
	l.log("ERROR", message, context)
}

// Debug logs a debug level message
func (l *Logger) Debug(message string, context map[string]interface{}) {
	l.log("DEBUG", message, context)
}

// log is a private method to handle the formatting of log entries
func (l *Logger) log(level, message string, context map[string]interface{}) {
	// Create timestamp
	timestamp := time.Now().Format(time.RFC3339)
	logEntry := fmt.Sprintf("[%s] [%s] %s", timestamp, level, message)

	// Log the context if available
	if context != nil {
		logEntry += fmt.Sprintf(" - Context: %+v", context)
	}

	// Print the log entry
	l.logger.Println(logEntry)
}
