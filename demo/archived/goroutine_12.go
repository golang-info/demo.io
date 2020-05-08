package main

import "fmt"

func fun(ch chan string) {
	for v := 0; v < 4; v++ {
		ch <- "hello ch"
	}
	close(ch)
}

func main() {
	c := make(chan string)

	go fun(c)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel closed ", ok)
			break
		}
		fmt.Println("Channel open ", res, ok)
	}
}