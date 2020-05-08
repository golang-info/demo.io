package main

import (
	"fmt"
)

func f(ch chan int) {
	fmt.Println("f function")
	ch <- 1
}

func main() {
	ch := make(chan int)
	go f(ch)
	//time.Sleep(1 * time.Second)
	<- ch
	fmt.Println("main function")
}
