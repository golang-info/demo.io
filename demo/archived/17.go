package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9
	ch <- 10

	for i := 0; i < 10; i++ {
		select {
		case <-ch:
			fmt.Println("random 1")
		case <-ch:
			fmt.Println("random 2")
		default:
			fmt.Println("error")
		}
	}
}
