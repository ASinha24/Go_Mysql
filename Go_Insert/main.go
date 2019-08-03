package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	fmt.Println("Go MySQL program")

	db, err := sql.Open("mysql", "root:system@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	fmt.Println("successfully connected to mysql db")

	insert, err := db.Query("INSERT INTO gouser VALUES('Thanos')")

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	fmt.Println("successfully inserted into user table")
}
