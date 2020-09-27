package main

import "fmt"

func main(){
	c:= make(chan int)
	go fibonacci(10, c)
	for i:= 0; i<12; i++ {
		foo, ok := <- c 
		if !ok {
			fmt.Println("done")
			return
		}
		fmt.Println(foo)
	}

}

func fibonacci( n int, c chan int){
	x, y := 0, 1
	for i := 0 ; i<n ;i++ {
		c <- x
		x,y = y,x+y
	}
	close(c)
}