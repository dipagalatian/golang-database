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
	script := "INSERT INTO customer(id, name) VALUES('mamat', 'Mamat');"

	// ExecContext can be used to send sql command that dont return rows (like insert, update, delete)
	_, err  := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert data customer")
	
}

func TestSelectSql(t *testing.T) {

	db := GetConnectionDB()
	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer;"
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("id:", id)
		fmt.Println("name:", name)
		
	}
}