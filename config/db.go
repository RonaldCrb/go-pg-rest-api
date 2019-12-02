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

	// create required tables
	offersTable := `
		CREATE TABLE btcoffers (
		  ID 					SERIAL PRIMARY KEY NOT NULL,
		  Title 			VARCHAR(255) NOT NULL,
			Trader 			VARCHAR(255) NOT NULL,
			Bank 				VARCHAR(255) NOT NULL,
			Currency 		VARCHAR(255) NOT NULL,
			Reputation	SMALLINT,
			Price 			FLOAT,
			Min 				FLOAT,
			Max 				FLOAT,
			Index 			SMALLINT NOT NULL,
		  CreatedAt   TIMESTAMP NOT NULL DEFAULT NOW(),
		  UpdatedAt   TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`

	usersTable := `
		CREATE TABLE users (
			id       	  SERIAL PRIMARY KEY NOT NULL,
			firstName   VARCHAR(255) NOT NULL,
			LastName 	  VARCHAR(255) NOT NULL,
			email		  	VARCHAR(255) NOT NULL,
  		createdAt   TIMESTAMP NOT NULL DEFAULT NOW(),
  		updatedAt   TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`

	_, err = DB.Exec(offersTable)
	if err != nil {
		log.Printf("[WARNING - CONFIG - DB] => %v", err)
	}

	_, err = DB.Exec(usersTable)
	if err != nil {
		log.Printf("[WARNING - CONFIG - DB] => %v", err)
	}

}
