# Backend-Gin-Gorm 
Backend for user grid.

### DB configs
Database name:MySQL
Installation 
https://www.mysql.com/downloads/

CREATE DATABASE users;


### .ENV Configs
port=3005 // backend run port

DB_URL = "username:password@tcp(localhost:3306)/users?charset=utf8mb4&parseTime=True&loc=Local"

### MIGRATE
go run migrate/migrate.go


### RUN
go run main.go


