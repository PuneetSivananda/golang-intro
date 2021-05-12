package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var con *sql.DB

func Connect() *sql.DB{
	// db, err:= sql.Open("sqlite3", "../todo.db")
	db, err:= sql.Open("sqlite3", "../todo.db")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to the database !...")
	con = db
	return con
}