package entity

// User represents a user in the system
// @Description User entity with basic information
type User struct {
	BaseModel
	// Username of the user
	Username string
	// Email address of the user
	Email string
	// Password (hashed, not exposed in JSON)
	Password string
}
