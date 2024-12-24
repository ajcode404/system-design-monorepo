package main

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "password"
	hostname = "localhost:3306"
	dbname   = "connection-pooling"
)

func main() {
	sql.Register("mysql", &mysql.MySQLDriver{})

}
