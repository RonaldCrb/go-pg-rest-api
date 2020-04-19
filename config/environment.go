package config

import "os"

// PgURL is the connection string for the database
var PgURL string = os.Getenv("CERBERUS_PG_URI")

// JWTSecret is the secret used to encrypt user passwords in the database
var JWTSecret string = os.Getenv("JWT_SECRET")

/*
// os.Getenv()
// gets an environment variable. It’s not possible to determine not set or empty. Use os.LookupEnv() to be able to do that.
// name := os.Getenv("NAME")

// os.Setenv() sets an environment variable.
// os.Setenv("NAME", "Flavio")

// os.Unsetenv() unsets an environment variable.
// os.Unsetenv("NAME")

// os.Clearenv() unsets all environment variables.
// os.Clearenv()

// os.Environ() returns a slice of strings containing all the environment variables in key=value format.
// vars := os.Environ()

// os.ExpandEnv() given a string, expands $VAR environment variables entries to the corresponding value.
// s := os.ExpandEnv("$NAME is italian")

// os.LookupEnv() returns the value of the environment variable in its first parameter if set, otherwise the second parameter is false. Allows to distinguish unset from empty value.
// name, ok := os.LookupEnv("NAME")
*/
