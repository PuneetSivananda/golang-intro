package main

import (
	"fmt"
	"strings"
)

/*
All strings are Mutable in go
*/
func main(){
	m1 := "My Name"

	fmt.Println(strings.Split(m1, " "))
}