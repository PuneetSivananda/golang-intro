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

	// Looping
	for i:= 0; i<10; i++{
		fmt.Println(i)
	}
	// infinite loop
	// for {
	// 	fmt.Println("Text")
	// }

	name:=[]string{"Rob", "Van", "Mike"}
	for item:= range(name){
		fmt.Println(item, name[item])
	}


	//maps

	mymap := make(map[string]interface{})
	mymap["name"] = "Testing"
	mymap["age"] = 20

	for k, v := range(mymap){
		fmt.Printf("Key is: %s and value is: %v", k, v)
		fmt.Println()
	}
}
