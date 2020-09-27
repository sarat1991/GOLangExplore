package main

import "fmt"


func greet1(ch chan string){
	 ch <-"Hello World 1"
}

func greet2(ch chan string){
	ch <- "Hello World 2"
}
func main(){
	ch1:= make(chan string)
	ch2:= make(chan string)
	go greet1(ch1)
	go greet2(ch2)
	for i:=0; i<2; i++ {
	select {
	case msg1:= <-ch1:{
		fmt.Println(msg1)
	}
	case msg2:= <-ch2: {
	fmt.Println(msg2)
	}
	}
}
}
