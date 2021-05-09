package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}){
	//Switch case implemented below in a indefinate loop
	for {
		time.Sleep(time.Second)
		select{
		case <-c:
			fmt.Println("Recieved")
		case <-quits:
			fmt.Println("Exiting...")
			os.Exit(0)
		}
	}
}

func main(){
	fmt.Println("Select is for Channels and is Async")
	
	c:= make(chan int, 2)
	quits:= make(chan struct{})

	go func(){
		c<- 1
		c<- 2
		c<- 233
		quits<- struct{}{}
	}()
	go Select(c, quits)	
	select{}
}