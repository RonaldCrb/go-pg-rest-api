package config

import (
	"database/sql"
	"fmt"

	// db driver
	_ "github.com/lib/pq"
)

// stablish database connection config constants
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

// DB variable of type pointer to DB type in sql package
var DB *sql.DB

func init() {

	// creates a string with all the db constants
	pgConfig := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// opens a connection with the database
	DB, err := sql.Open("postgres", pgConfig)

	// checks for connection errors
	if err != nil {
		panic(err)
	}

	// defers closing the db connection
	defer DB.Close()

	// ping database
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	// logs on connection success
	fmt.Println("PostgreSQL Database ready to GO!")
}
