package main

import (
	"context"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "salah"

	script := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1";
	
	rows, err := db.QueryContext(ctx, script, username, password)

	if err != nil
	defer rows.Close()

	if rows.Next() {
		var username string
		err := rows.Scan(&username)

		if err !== nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login", username)		
	}
}

func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "irwan'; DROP TABLE user; #"
	password := "irwan"

	script := "INSERT INTO user(username, password) VALUES(?, ?)";
	
	_, err := db.ExecContext(ctx, script, username, password)

	if err != nil
	
	fmt.Println("Success insert new user")		
}