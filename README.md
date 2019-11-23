# GO and PostgreSQL JSON REST API

this is a sample RESTful API written in go.

## launching a containerized database for testing purposes

First you need to Install Docker

then run this command to run a postgres container which maps port 5432 to localhost:5432

```shell
docker run --name pg -p 5432:5432 -d postgres
```

then this command will give you a psql shell with the appropriate user:

```shell
docker exec -it pg psql postgres postgres
```

## creating databases and tables

inside the psql shell, create a database with this command (dont forget semi-colons)

```sql
CREATE DATABASE users;
```

after database is created, run this command to create a users table

```sql
CREATE TABLE users (
	id       	  SERIAL PRIMARY KEY NOT NULL,
	firstName     	  VARCHAR(255) NOT NULL,
	LastName 	  VARCHAR(255) NOT NULL,
	email		  VARCHAR(255) NOT NULL,
        createdAt         TIMESTAMP NOT NULL DEFAULT NOW(),
        updatedAt         TIMESTAMP NOT NULL DEFAULT NOW()
);
```
## endpoints

You can execute the binary file directly by typing
```shell
./main
```

this will make available the following endpoints:

### GET=> http://localhost:8080/home
returns a JSON object with:
a greeting message

### GET => http://localhost:8080/healthcheck
returns a JSON object with:
a healthcheck status message and timestamp

### POST => http://localhost:8080/users
returns a JSON object with:
a confirmation message and status code

### GET => http://localhost:8080/users
returns a JSON array with:
user objects 

### GET => http://localhost:8080/users/:id
returns a JSON object with:
user properties

### PUT => http://localhost:8080/users/:id
returns a JSON object with:
a confirmation message and status code

### DELETE => http://localhost:8080/users/:id
returns a JSON object with:
a confirmation message and status code
