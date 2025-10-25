package main

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestAutoIncrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "eko@gmail.com"
	comment := "Test Komen"

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)";

	result, err := db.ExecContext(ctx, script, email, comment)

	if err != nil {
		panic(err)
	}

	insertId, err := result.LastInsertId()

	if err != nil {
		panic(err)
	}
	
	fmt.Println("Success insert new comment with id", insertId)		
}