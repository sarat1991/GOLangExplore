package main

import "fmt"

func main() {
	var fruits [2]string

	fruits[0] = "Apple"
	fruits[1] = "Orange"

	fmt.Println(fruits)
	fmt.Println(fruits[1])

	cars := [2]string{
		"kia","tata"}

	fmt.Println(cars)
	fmt.Println(cars[0])

	carsSlice :=[] string{"kia", "tata", "hyundai", "maruthi"}

	fmt.Println(len(carsSlice))
	fmt.Println(carsSlice[0:5])
}