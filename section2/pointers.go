package main

import "fmt"

//Pass by ref uses pointers
func swap(m1, m2 * int){
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp

}

func main() {
	m1, m2 := 2, 3
	ptr := &m1
	// & referece operator
	// * derefence operator
	fmt.Println(ptr)
	fmt.Println(*ptr)

	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)

}