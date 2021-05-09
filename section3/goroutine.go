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

func superHeavy(){
	for{
		fmt.Println("Super Heavy function")
		time.Sleep(time.Second * 2)
	}

}

func main(){
	// heavy() // normal function
	go heavy() //goroutine
	go superHeavy()
	fmt.Println("VSCODE go")
	select{}
}