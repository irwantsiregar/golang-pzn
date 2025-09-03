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

func TestQuerySql(t *testing.T) {
	db := GetConnection()

	defer db.Close()

	ctx := context.Background()

	script := "SELECT id, name FROM customer";
	
	rows, err := db.QueryContext(ctx, script)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string

		err = rows.Scan(&id, &name)

		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}

	fmt.Println("Success get customer")
}