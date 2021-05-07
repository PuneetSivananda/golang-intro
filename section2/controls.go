package main

import "fmt"

func main(){
	fmt.Println("Foo bar")
	fmt.Println("Hello World")

	flag:= true
	
	if flag && 2 ==5 || 4 ==4{
		fmt.Println("True")
	}	else{
		fmt.Println("False")
	}

	f := &flag
	
	if *f{
		fmt.Println("True")
	} else if !*f {
		fmt.Println("False")
	}

	if f == nil {
		fmt.Println("Value is nil")
	} else if *f {
		fmt.Println("True")
	}	else{
		fmt.Println("False")
	}

	
}
