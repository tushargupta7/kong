package repositories

import (
	"database/sql"
	"github.com/tushargupta7/kong/database"
	"github.com/tushargupta7/kong/dtos"
	"github.com/tushargupta7/kong/errors"
)

// GetUserByUsername fetches a user record from the database by their username
func GetUserByUsername(username string) (*dtos.User, error) {
	var user dtos.User

	// Query to fetch user by username
	query := "SELECT username, password, role FROM users WHERE username = $1"
	err := database.DB.QueryRow(query, username).Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New(404, "User not found", err, nil)
		}
		return nil, errors.New(500, "Failed to fetch user", err, nil)
	}

	return &user, nil
}
