package users

import (
	"errors"
	"log"
	"time"

	"github.com/RonaldCrb/go-mc/config"
)

// User represents a user instance
type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// CreateUser creates a new user instance
func (u User) CreateUser() error {
	usr := `INSERT INTO 
						users (firstName, lastName, email)
						VALUES ($1, $2, $3)`

	stmt, err := config.DB.Prepare(usr)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rows, err := stmt.Exec(u.FirstName, u.LastName, u.Email)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := rows.RowsAffected()

	if aff != 1 {
		err = errors.New("[ERROR] => More than 1 rows where affected")
	}

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// AllUsers returns a slice of User (all users in users table)
func AllUsers() ([]User, error) {
	q := "SELECT * FROM users ORDER BY ID ASC"

	rows, err := config.DB.Query(q)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	usrs := make([]User, 0)

	for rows.Next() {
		usr := User{}
		err := rows.Scan(&usr.ID, &usr.FirstName, &usr.LastName, &usr.Email, &usr.CreatedAt, &usr.UpdatedAt)

		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		usrs = append(usrs, usr)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return usrs, nil
}

// FindUser returns a user instance from the database
func (u User) FindUser() (User, error) {
	q := "SELECT * FROM users WHERE id = $1"

	row := config.DB.QueryRow(q, u.ID)

	usr := User{}

	err := row.Scan(&usr.ID, &usr.FirstName, &usr.LastName, &usr.Email, &usr.CreatedAt, &usr.UpdatedAt)
	if err != nil {
		log.Fatal(err)
		return User{}, err
	}

	return usr, nil
}

// UpdateUser updates the data for a user instance in the database
func (u User) UpdateUser() error {
	q := "UPDATE users SET FirstName=$1, LastName=$2, Email=$3, UpdatedAt=now() WHERE ID = $4"

	if u.FirstName == "" || u.LastName == "" || u.Email == "" {
		err := errors.New("[ERROR] => FirstName, LastName and Email fields are required")
		return err
	}

	stmt, err := config.DB.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	row, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	aff, err := row.RowsAffected()
	if aff != 1 {
		err = errors.New("[ERROR] => More than 1 rows where affected")
	}

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// DeleteUser deletes a user instance from the database
func (u User) DeleteUser() error {
	q := `DELETE FROM users WHERE ID=$1`

	_, err := config.DB.Exec(q, u.ID)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
