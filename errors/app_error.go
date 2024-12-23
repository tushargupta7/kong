package errors

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/logger"
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
	log.Error(message, context)

	return &AppError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
		Context:    context,
	}
}

func ErrorHandler(c *fiber.Ctx) error {
	// Check if the error is of type AppError
	if err := c.Next(); err != nil {
		if appErr, ok := err.(*AppError); ok {
			// If it's an AppError, return the response based on AppError properties
			return c.Status(appErr.StatusCode).JSON(fiber.Map{
				"error":   appErr.Message,
				"details": appErr.Context,
			})
		}

		// If it's not an AppError, return a generic 500 internal server error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}
	return nil
}
