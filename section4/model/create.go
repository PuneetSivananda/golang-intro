package model

import (
	"database/sql"
	"fmt"	"log"
)


func CreateTODO() error {
	// tx, _ := con.Begin()
	result, err := con.Exec("insert into TODO(name, todo) values(?, ?)", "Puneet", "Todo1")
	fmt.Println(result)
	// tx.Commit()
	defer con.Close()
	if err != nil {
		return err
	}
	return nil
	// insertQ, err := con.Exec("INSERT INTO TODO VALUES(?, ?)", "Puneet", "Todo1")
	// defer con.Close()
	// if err != nil {
	// 	fmt.Println(insertQ)
	// 	return err
	// }
	// fmt.Println(insertQ)
	// return nil
}

func insertStudent(db *sql.DB, code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement. This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
}