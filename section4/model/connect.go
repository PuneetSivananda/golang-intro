package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var con *sql.DB

func CreateTable(db *sql.DB) {
	createTODOTableSQL := `CREATE TABLE TODO (
		"name" TEXT,		
		"todo" TEXT
	  );`

	statement, err := db.Prepare(createTODOTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("TODO table created")
}

func CreateDB(){
	os.Remove("todo.db")

	log.Println("Creating todo.db...")
	file, err := os.Create("todo.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("todo.db created")
}

func Connect() *sql.DB{
	// CreateDB()
	// CreateTable(db)
	db, err:= sql.Open("sqlite3", "./todo.db")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Connected to the database !...")
	con = db
	return con
}