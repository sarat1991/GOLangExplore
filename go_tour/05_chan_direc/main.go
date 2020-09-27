package main

import "fmt"

func ping (ping chan<- string, message string){
	ping <- message
}

func pong(ping <-chan string, pong chan <- string){
	message := <-ping
	pong <- message
}

func main () {
	pg := make(chan string, 1)
	pog := make(chan string, 1)
	ping(pg, "Hello World")
	pong(pg, pog)
	fmt.Println(<-pog)
}