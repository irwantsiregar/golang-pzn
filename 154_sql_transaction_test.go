package main

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)
func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	
	tx, err := db.Begin()
	
	if err != nil {
		panic(err)
	}
	
	script := "INSERT INTO comments(email, comment) VALUES(?, ?)";

	for i := 0; i < 10; i++ {
		email := "irwan" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke" + strconv.Itoa(i)


		// result, err := statement.QueryContext(ctx, email, comment)
		result, err := tx.ExecContext(ctx, script, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)		
	}


	err = tx.Commit()

	if err != nil {
		panic(err)
	}
}