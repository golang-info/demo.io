package main

import (
	"fmt"
)

func main() {
	done := make(chan bool)

	go func() {
		fmt.Println("Goroutine message")
		done <- true
	}()

	fmt.Println("Main message")
	<-done
}
