package main

// Buffered channels
import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Buffered Channels")
	c:= make(chan int, 3)
	go func(){
		c<- 1
		c<- 2
		c<- 3
		c<- 4
		close(c)
	}()
	
	time.Sleep(time.Second *1)

	for i:= range(c){
		fmt.Println(i)
	}
	
}