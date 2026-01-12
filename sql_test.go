package golangdatabase

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
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

func TestQuerySqlComplex(t *testing.T) { 

	db := GetConnectionDB()
	defer db.Close()

	ctx := context.Background()
	
	script := "SELECT id, name, email, balance,rating, birth_date, created_at, married FROM customer;"
	rows, err  := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {

		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate sql.NullTime 
		var createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &createdAt, &married)
		if err != nil {
			panic(err)
		}

		fmt.Println("========================")
		fmt.Println("id:", id)
		fmt.Println("name:", name)
		if email.Valid {
			fmt.Println("email:", email.String)
		}
		fmt.Println("balance:", balance)
		fmt.Println("rating:", rating)
		fmt.Println("createdAt:", createdAt)
		if birthDate.Valid {
			fmt.Println("birthDate:", birthDate.Time)
		}
		fmt.Println("married:", married)
		fmt.Println("========================")
	}
	
}

func TestSqlInjection(t *testing.T) {

	db := GetConnectionDB()
	defer db.Close()

	ctx := context.Background()

	// simulate user input for login
	// username := "admin"
	// password := "admin" 

	// simulate sql injection attack
	// explanation: by inputting this value for username
	// the sql script will be manipulated to always return true
	// because the part after # will be ignored by mysql (# is comment in mysql)
	// so the script will be like this:
	// SELECT username FROM user WHERE username = 'admin' # ' AND password = 'asdf' LIMIT 1;
	// which is equivalent to:
	// SELECT username FROM user WHERE username = 'admin';
	// thus the attacker can login without knowing the correct password
	username := "admin'; #"
	password := "asdf"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1;"
	fmt.Println("script login:", script)
	
	rows, err  := db.QueryContext(ctx, script)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		
		fmt.Println("Login success:", username)
	} else {
		fmt.Println("Login Failed")
	}
	
}