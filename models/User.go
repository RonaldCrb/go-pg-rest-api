package models

// User represents a user instance
type User struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	// Background *Background `json:"background,omitempty"`
}
