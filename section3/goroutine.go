package main

import (
	"fmt"
	"time"
)


func heavy(){
	time.Sleep(time.Second *5)
	fmt.Println("Heavy function")
}

func main(){
	// heavy() // normal function
	go heavy() //goroutine
	fmt.Println("VSCODE go")
}