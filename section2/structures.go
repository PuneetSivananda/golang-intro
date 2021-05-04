package main

import "fmt"

type Car struct{
	Name string
	Age int
	ModelNo int
}

func (c Car) Print(){
	fmt.Println(c)
}

func (c Car) Drive(){
	fmt.Println("Driving...", c.Name)
}

func (c Car) GetName() string{
	return c.Name
}


func main() {
	fmt.Println("Hello World")
	c:= Car{
		Name:"Indica",
		Age: 10,
		ModelNo: 1100,
	}
	c.Print()
	c.Drive()
	c.GetName()
}