package main

import "fmt"

func main() {
	p := 10
	pointerP := &p
	
	fmt.Println(pointerP)
	fmt.Println(*pointerP)

	*pointerP = 25

	fmt.Println(p)

}