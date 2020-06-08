package main

import "fmt"

func main() {
	c := make(chan int)
	q := 10
	go generateInts(c, q)

	for i := range c {
		fmt.Println(i)
	}
}

func generateInts(c chan int, q int) {
	for i := 0; i <q; i++ {
		c <- i
	}
	close(c)
}
