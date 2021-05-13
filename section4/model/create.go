package model

import (
	"fmt"
	"log"
)


func CreateTODO(name, todo string) error {
	fmt.Println("--->", name, todo)
	_, err := con.Exec("insert into TODO(name, todo) values(?, ?)", name, todo)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func DeleteByName(name string) error {
	insertQ, err := con.Query("DELETE FROM TODO WHERE name=?", name)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer insertQ.Close()
	return nil
}

