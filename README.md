# GO-MC (Model, Controller) REST API


## creating databases
```sql
CREATE DATABASE users;
```
## creating table
```sql
CREATE TABLE users (
	id       		SERIAL PRIMARY KEY NOT NULL,
	firstName 	VARCHAR(255) NOT NULL,
	LastName 	VARCHAR(255) NOT NULL,
	email				VARCHAR(255) NOT NULL,
  createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
  updatedAt TIMESTAMP NOT NULL DEFAULT NOW()
);
```
## inserting records (rows)
```sql
INSERT INTO users 
(FirstName, LastName, Email) 
VALUES 
('rick', 'grimes', 'rick.grimes@email.com'),
('mary', 'jane', 'mary.jane@email.com'),
('luisa', 'diaz', 'luisa.diaz@email.com'),
('johana', 'mata', 'johana.mata@email.com'),
('juan', 'perez', 'juan.perez@email.com');
```
