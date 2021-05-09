package main

// lambda functions
import (
	"fmt"
)

func main(){
	go func(){
		fmt.Println("Lambda functions : 1")
	}()
	go func(){
		fmt.Println("Lambda functions : 2")
	}()
	fmt.Println("End")
	select{} // This causes deadlocks avoid them

}
