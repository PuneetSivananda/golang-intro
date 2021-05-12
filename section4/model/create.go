package model

import "fmt"

func CreateTODO() error {
	insertQ, err := con.Exec("INSERT INTO TODO VALUES(?, ?)", "Puneet", "Todo1")
	defer con.Close()
	if err != nil {
		fmt.Println(insertQ)
		return err
	}
	return nil
}