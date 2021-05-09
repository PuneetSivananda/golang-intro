package main

import (
	"fmt"
	"time"
)

// The Pinger print a ping and waits for a pong
func pinger(pinger <-chan int, ponger chan<- int){
	for{
		<- pinger
		fmt.Println("PING")
		time.Sleep(time.Second *1)
		ponger<- 1
	}
}

// The Ponger print a pong and waits for a ping
func ponger(pinger chan<- int, ponger <-chan int){
	for{
		<-ponger
		fmt.Println("PONG")
		time.Sleep(time.Second *1)
		pinger<- 1
	}
}


func main(){
	fmt.Println("Ping Ping in GO")
		
	ping:= make(chan int)
	pong:= make(chan int)

	go pinger(ping, pong)
	go ponger(ping, pong)
	
	// Initialize the ping 
	ping<-1
	// for{
	// 	time.Sleep(time.Second)
	// }

	// indefinate loop
	select{}

}
