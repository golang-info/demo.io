package main

import "fmt"

func main() {
	ch := make(chan int)

	for i := 0 ; i < 5 ; i++ {
		go func(s int) {
			fmt.Println("hello", <-ch)
		}(i)
		ch <- i
	}

}