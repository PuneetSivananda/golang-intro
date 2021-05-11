package model

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB{
	db, err:= sql.Open("sqlite3", ":memory:")
	if err != nil{
		log.Fatal(err)
	}
	return db
}