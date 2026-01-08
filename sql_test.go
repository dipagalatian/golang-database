package golangdatabase

import (
	"context"
	"fmt"
	"testing"
)


func TestExecSql(t *testing.T) {

	// call function GetConnectionDB from database.go
	// to get *sql.DB connection
	// remember to close connection after use
	// we dont need to open connection again because
	// GetConnectionDB already handle it
	db := GetConnectionDB()
	defer db.Close()

	// ctx can be used to set timeout or cancel sinyal
	ctx := context.Background()

	// example insert data sql
	script := "INSERT INTO customer(id, name) VALUES('joko', 'Joko');"

	// ExecContext can be used to send sql command that dont return rows (like insert, update, delete)
	_, err  := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data customer")
	
}