package main

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	script := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1";
	
	rows, err := db.QueryContext(ctx, script)

	if err != nil
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err !== nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	}
	else {
				fmt.Println("Gagal Login", username)		
	}
}