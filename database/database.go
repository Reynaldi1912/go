package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	port := "3306"
	username := "root"
	host := "localhost"
	database := "kasir"
	name := "mysql"

	connection := username + "@tcp(" + host + ":" + port + ")/" + database
	db, err := sql.Open(name, connection)

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	return db
}
