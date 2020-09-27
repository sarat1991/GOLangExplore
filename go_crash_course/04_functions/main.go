package main

import "fmt"

func main () {
	fmt.Println(greeting("Sarath"))
	fmt.Println(sum(1,2))
}

func greeting (name string) string{
	return "Hello "+name
}

func sum( a int , b int) int {
	return a+b
}