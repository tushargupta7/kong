package errors

import (
	"fmt"
	"github.com/tushargupta7/kong/logger" // Import your logger package
)

// AppError defines a custom error structure
type AppError struct {
	StatusCode int                    `json:"-"`
	Message    string                 `json:"message"`
	Err        error                  `json:"-"`
	Context    map[string]interface{} `json:"context"`
}

var log = logger.NewLogger()

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("message: %s, error: %v, context: %v", e.Message, e.Err, e.Context)
	}
	return fmt.Sprintf("message: %s, context: %v", e.Message, e.Context)
}

func New(statusCode int, message string, err error, context map[string]interface{}) *AppError {
	// Log the error when creating an AppError
	log.Error(message, context)

	// Return the AppError
	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
		Context:    context,
	}
}
