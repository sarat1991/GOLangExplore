package main

import ("fmt" 
"time")

func fib(ch chan bool){
	fmt.Print("Working....")
	time.Sleep(time.Second)
	fmt.Println("Done")
	ch <- true
}
func main (){
	ch := make(chan bool)
	go fib(ch)
	<-ch
}

