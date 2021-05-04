package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

type Lambo struct{
	LamboModel string
}

func NewModel(arg string) Car{
	return &Lambo{arg}
}

type Chevy struct{
	ChevyModel string
}

func (l *Lambo) Stop(){
	fmt.Println("Stopping Lambo")
}

func (l *Lambo) Drive(){
	fmt.Println("Driving the Lamborghini")
	fmt.Println(l.LamboModel)
}


func (c *Chevy) Drive(){
	fmt.Println("Driving the Chevy")
	fmt.Println(c.ChevyModel)
}


func main(){
	l:= NewModel("Gallardo")
	c:= Chevy{"Bolt"}
	l.Drive()
	c.Drive()
}