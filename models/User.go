package models

import (
	"github.com/RonaldCrb/go-mc/config"
)

// User represents a user instance
type User struct {
	id        int32
	firstName string
	lastName  string
	email     string
}

// AllUsers returns a slice of User (all users in users table)
func AllUsers() ([]User, error) {
	rows, err := config.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	usrs := make([]User, 0)

	for rows.Next() {
		usr := User{}
		err := rows.Scan(&usr.id, &usr.firstName, &usr.lastName, &usr.email)
		if err != nil {
			return nil, err
		}
		usrs = append(usrs, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return usrs, nil
}
