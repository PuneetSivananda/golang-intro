package main

// Buffered channels
import (
	"fmt"
	"time"
)
type Car struct{
	Model string
}

func main(){
	fmt.Println("Buffered Channels")
	c:= make(chan Car, 3)
	go func(){
		c<- Car{"1"}
		c<- Car{"2"}
		c<- Car{"3"}
		c<- Car{"4"}
		close(c)
	}()

	time.Sleep(time.Second *1)

	for i:= range(c){
		fmt.Println(i.Model)
	}
	
}