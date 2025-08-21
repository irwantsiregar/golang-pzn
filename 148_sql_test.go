package main

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestExectSql(t *testing.T) {
	db := GetConnection()

	fmt.Println("This GetConnection")

	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO customer (id, name) VALUES ('101', 'Irwan')";
	
	_, err := db.ExecContext(ctx, script)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new customer")
}