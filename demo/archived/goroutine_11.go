package main

//Go program to illustrate send and receive operation

import "fmt"

func fun(ch chan int) {
	fmt.Println(234 + <-ch)
}

func main() {

	ch := make(chan int)
	go fun(ch)
	ch <- 23
	
}