# GO-MC (Model, Controller) REST API


## creating databases
```sql
CREATE DATABASE users;
```
## creating table
```sql
CREATE TABLE users (
	id       		char(14) PRIMARY KEY NOT NULL,
	first_name 	varchar(255) NOT NULL,
	last_name 	varchar(255) NOT NULL,
	email				varchar(255) NOT NULL
);
```
## inserting records (rows)
```sql
INSERT INTO users 
(id, first_name, last_name, email) 
VALUES 
('1', 'ron', 'alonzo', 'alonzo.ronald@gmail.com'),
('2', 'maredy', 'glines', 'maredy.glines@gmail.com'),
('3', 'luisa', 'diaz', 'luisa.diaz@gmail.com'),
('4', 'johana', 'mata', 'johana.mata@gmail.com'),
('5', 'juan', 'perez', 'juan.perez@gmail.com');```
```
