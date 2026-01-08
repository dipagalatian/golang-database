package golangdatabase

import (
	"database/sql"
	"time"
)

func GetConnectionDB() *sql.DB {
	dsn := "root:password1234@tcp(127.0.0.1:3306)/db_belajar"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10) // set the maximum number of connections in the idle connection pool
	db.SetMaxOpenConns(100) // set the maximum number of open connections to the database
	db.SetConnMaxIdleTime(5 * time.Minute) // set the maximum amount of time a connection may be idle
	db.SetConnMaxLifetime(60 * time.Minute) // set the maximum amount of time a connection may be reused

	return db
}