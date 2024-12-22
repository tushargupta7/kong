package dtos

// LoginRequest represents the structure of the login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
