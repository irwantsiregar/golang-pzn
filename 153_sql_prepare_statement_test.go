package main

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)
func TestExecSqlParameter(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	script := "INSERT INTO comments(email, comment) VALUES(?, ?)";

	statement, err := db.PrepareContext(ctx, script)
	
	if err != nil {
		panic(err)
	}

	defer statement.Close()

	for i := 0; i < 10; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Komentar ke" + strconv.Itoa(i)


		// result, err := statement.QueryContext(ctx, email, comment)
		result, err := statement.ExecContext(ctx, email, comment)

		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id ", id)		
	}
}