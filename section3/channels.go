package main

// lambda functions
import (
	"fmt"
	"time"
)

func main(){
	// var c1 chan int // creates a channel
	c:= make(chan int) // makes a channel
	fmt.Println("Producers and Consumers")
	//send in a go routine
	go func(){
		c <- 143
	}()
	//sniff the channel
	val:= <-c
	fmt.Println(val)
	fmt.Println(c)
	time.Sleep(time.Second *1)
	go func(){
		c <- 52
	}()

	val= <-c
	fmt.Println(val)
	
	fmt.Println(c)
}