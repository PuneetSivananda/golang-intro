package main

import "fmt"

func main() {
	fmt.Println("Switch Case")
	//Continue, break, switch

	for i:=0; i<10; i++{
		if i%2==0{
			continue
		}
		fmt.Println(i)
	}

	for i:=1; i<10; i++{
		if i%2==0{
			break
		}
		fmt.Println(i)
	}

	day:= "Fri"

	switch day {
	case "Fri":
		fmt.Println("TGIF")
	case "Mon", "Tue", "Wed", "Thu":
		fmt.Println("Boring")
	}	
}