package main

// lambda functions
import (
	"fmt"
)

func main(){
	var c1 chan int // creates a channel
	c2:= make(chan int) // makes a channel
	fmt.Println("Producers and Consumers")
	fmt.Println(c1)
	fmt.Println(c2)
}