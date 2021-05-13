package model

import (
	"log"
)


func CreateTODO(name, todo string) error {
	createQ, err := con.Query("insert into TODO(name, todo) values(?, ?)", name, todo)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer createQ.Close()
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

