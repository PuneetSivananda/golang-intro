package main

import (
	"fmt"
	"time"
)


func heavy(){
	for{
		fmt.Println("Heavy function")
		time.Sleep(time.Second * 1)
	}
}

func main(){
	// heavy() // normal function
	go heavy() //goroutine
	fmt.Println("VSCODE go")
	time.Sleep(time.Second * 5)
}