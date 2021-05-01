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
	m2 := "is Slim Shady"
	fmt.Println(strings.Split(m1, " "), m1+m2)
}