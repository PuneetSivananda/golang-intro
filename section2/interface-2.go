package main

import "fmt"

func Anything(anything interface{}){
	fmt.Println(anything)
}

func main(){
	fmt.Println("Interfaces contd")
	Anything(2.4)
	Anything("Hello")
	Anything(struct{}{})

	mymap := make(map[string]interface{})
	mymap["name"] = "SNAFU"
	mymap["age"] = 10
	fmt.Println(mymap)
}