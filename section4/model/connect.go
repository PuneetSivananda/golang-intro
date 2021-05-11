package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

vat con *sql.DB

func Connect() *sql.DB{
	db, err:= sql.Open("sqlite3", ":memory:")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to the database !...")
	con = db
	return con
}