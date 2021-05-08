package main

import (
	"fmt"
)

func main(){
	var ( 
		m1 = 99
		m2 = 1024
	)
	var m3 int32
	var m4 int64

	m3 = 22
	m4 = 45

	m5:= 2
	m6:= 4
	fmt.Println(m1 + m2)
	// fmt.Println(m3 + m4)
	fmt.Println(int64(m3) + m4)
	fmt.Println(m5, m6)
}