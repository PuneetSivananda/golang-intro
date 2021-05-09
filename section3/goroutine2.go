package main

// lambda functions
import (
	"fmt"
	"sync"
)

func main(){
	//Using the Sync Package
	// wait groups
	wg:= &sync.WaitGroup{}
	// add, done and wait
	
	// Mutexes (critical resources)
	// mut:= &sync.Mutex{}
	// mut.Lock() 
	// mut.Unlock()

	wg.Add(2)
	go func(){
		fmt.Println("Lambda functions : 1")
		wg.Done()
	}()
	go func(){
		fmt.Println("Lambda functions : 2")
		wg.Done()
	}()
	
	wg.Wait()
	fmt.Println("End")
	// select{} // This causes deadlocks avoid them
}
