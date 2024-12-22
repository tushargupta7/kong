package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/tushargupta7/kong/errors"
	"github.com/tushargupta7/kong/repositories"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// To be kept as env variable or fetch from secure key store
var jwtSecretKey = []byte("dummy_secret")

// ValidatePassword compares the provided password with the stored hashed password
func ValidatePassword(storedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}

// GenerateJWT generates a JWT token for the authenticated user
func GenerateJWT(username, role string) (string, error) {
	claims := &jwt.MapClaims{
		"username": username,
		"role":     role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

// Login handles the login process, validates user credentials, and returns a JWT token
func Login(username, password string) (string, error) {
	user, err := repositories.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	if !ValidatePassword(user.Password, password) {
		return "", errors.New(401, "Invalid username or password", nil, nil)
	}

	token, err := GenerateJWT(user.Username, user.Role)
	if err != nil {
		return "", errors.New(500, "Failed to generate token", err, nil)
	}

	return token, nil
}
