package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)


func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "developer:@tcp(localhost:3306)/belajar_golang_database")

	fmt.Println("This GetConnection")
	
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
}