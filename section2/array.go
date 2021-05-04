package main

import "fmt"

func todo(){
	//var arr[] int
	arr1 := []int{2,4,6,8}
	arr2 := []string{"Hi","My","Name","is","Rudolph"}
	// append function adds an item to the array
	arr2 = append(arr2, "!!")
	fmt.Println(arr1, arr2)
}

func main(){
	fmt.Println("Arrays")
	todo()
}