package main

import "fmt"

func main () {

	x:= 12
	y:= 10

	if x < y {
		fmt.Println(x , " is less than " , y)
	}else {
		fmt.Println(x ," is greater than", y)
	}

	color := "green"

	switch color {
	case "green":
			fmt.Println("Color is green")
			break
	case "blue": 
			fmt.Println("Color is blue")
			break
	
	}
}