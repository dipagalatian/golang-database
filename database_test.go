package golangdatabase

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestEmpty(t *testing.T) {
	
}

func TestOpenConnection(t *testing.T) {
	// Format: username:password@tcp(host:port)/database_name
    dsn := "root:password1234@tcp(127.0.0.1:3306)/db_belajar"
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test access the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully open connection database")
	
}