package model

import (
	"database/sql"
	"fmt"
	"log"
)

func Connect() *sql.DB{
	db, err:= sql.Open("sqlite3", ":memory:")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to the database !...")
	return db
}