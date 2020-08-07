package main

/*
	https://dev.to/spindriftboi/concurrency-in-go-using-channels-and-handling-race-conditions-3n24
*/

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 29
	}()

	fmt.Println(<-c) // 29
}
