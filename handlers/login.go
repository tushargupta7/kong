package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tushargupta7/kong/auth"
	"github.com/tushargupta7/kong/dtos"
	"github.com/tushargupta7/kong/errors"
)

// Login endpoint to authenticate the user and generate a JWT token
func Login(c *fiber.Ctx) error {
	var req dtos.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return errors.New(fiber.StatusBadRequest, errors.ErrInvalidPayload, err, nil)
	}

	// Validate user credentials and generate JWT token
	token, err := auth.Login(req.Username, req.Password)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
