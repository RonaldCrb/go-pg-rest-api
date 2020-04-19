package config

import "os"

// PgURL is the connection string for the database
var PgURL string = os.Getenv("PG_CONNECTION")

// JWTSecret is the secret used to encrypt user passwords in the database
var JWTSecret string = os.Getenv("JWT_SECRET")
