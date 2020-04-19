package config

import (
	"database/sql"
	"fmt"
	"log"

	// db driver
	_ "github.com/lib/pq"
)

// DB variable of type pointer to DB type in sql package
var DB *sql.DB

func init() {

	var err error

	// opens a database connection using the URL from environment.go (PgURL)
	DB, err = sql.Open("postgres", PgURL)
	if err != nil {
		log.Printf("[warning] => %v", err)
	} else {
		fmt.Println("PostgreSQL Database ready to GO!")
	}

	usersTable := `
		CREATE TABLE users (
			id       	  SERIAL PRIMARY KEY NOT NULL,
			firstName   VARCHAR(255) NOT NULL,
			lastName 	  VARCHAR(255) NOT NULL,
			email		  	VARCHAR(255) NOT NULL,
			password  	VARCHAR(255) NOT NULL,
  		createdAt   TIMESTAMP NOT NULL DEFAULT NOW(),
  		updatedAt   TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`

	_, err = DB.Exec(usersTable)
	if err != nil {
		log.Printf("[WARNING - CONFIG - DB] => %v", err)
	}

}
