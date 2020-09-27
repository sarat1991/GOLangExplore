package main

import ("fmt"
"time")

func say(word string) {
	for i:= 0; i<10; i++ {
		time.Sleep(100* time.Millisecond)
		fmt.Println(word)
	}
}

func main() {
	go say("World!!")
	say("Hello")
}