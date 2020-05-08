package main

import "fmt"

func count(i int, ch chan int) {
	fmt.Println("Counting: ", i)
	ch <- i
}

func main() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go count(i, chs[i])
	}

	for _, ch := range chs {
		fmt.Println("Counting num: ", <-ch)
	}
}
