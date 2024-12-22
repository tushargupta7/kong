package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/tushargupta7/kong/errors"
	"strings"
)

// Define the secret key for JWT signing (same as the one used during generation)
var jwtSecretKey = []byte("dummy_secret")

// JWTMiddleware validates the JWT token and checks the role
func JWTMiddleware(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Extract the token from the Authorization header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return errors.New(401, "Authorization header missing", nil, nil)
		}

		// Extract the token from the header
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Check the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		})

		if err != nil {
			return errors.New(401, "Invalid token", err, nil)
		}

		// Validate the token
		claims, ok := token.Claims.(*jwt.MapClaims)
		if !ok || !token.Valid {
			return errors.New(401, "Invalid token", nil, nil)
		}

		// Check the role
		role := (*claims)["role"].(string)
		if requiredRole != "" && role != requiredRole {
			return errors.New(403, "Insufficient permissions", nil, nil)
		}

		// Store user info in context for future use
		c.Locals("username", (*claims)["username"].(string))
		c.Locals("role", role)

		// Proceed to the next handler
		return c.Next()
	}
}
