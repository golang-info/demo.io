package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3)
	q := 10
	go worker(c, 1)
	go worker(c, 2)
	go worker(c, 3)
	go generateInts(c, q)

	time.Sleep(10 * time.Second)
}

func worker(c chan int, id int) {
	for i := range c {
		fmt.Printf("worker %d received value %d\n", id, i)
		time.Sleep(1 * time.Second)
	}
}

func generateInts(c chan int, q int) {
	for i := 0; i < q; i++ {
		c <- i
		fmt.Printf("Sent value %d\n", i)
		time.Sleep(1 * time.Millisecond)
	}
	close(c)
}
