package dtos

// User represents the structure of a user record in the database
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"` // Role can be 'admin' or 'user'
}
