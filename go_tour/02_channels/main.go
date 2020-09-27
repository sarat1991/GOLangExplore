package main

import "fmt"

func sum(c chan int, s[] int) {
	sum := 0
	for _,v := range s {
		sum += v
	}
	c <- sum
}

func main () {
	input := []int{1,2, 3, 4, 5}
	c := make(chan int, 10)
	go sum(c, input[:len(input)/2])
	go sum(c, input[len(input)/2:])
	x, y := <-c, <-c
	fmt.Println(x, " ", y, " ", x+y)
	newInput := input[:len(input)/2]
	for _, v := range newInput{
		fmt.Println (v)
	}
}