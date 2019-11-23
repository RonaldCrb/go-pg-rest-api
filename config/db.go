package config

import (
	"database/sql"
	"fmt"

	// db driver
	_ "github.com/lib/pq"
)

// DB variable of type pointer to DB type in sql package
var DB *sql.DB

func init() {
	var err error
	// opens a connection with the database
	DB, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	// logs on connection success
	fmt.Println("PostgreSQL Database ready to GO!")
}
