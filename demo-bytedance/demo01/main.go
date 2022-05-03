package main

import "fmt"

func main() {
	ch := make(chan int, 0)

	ch <- 666
	x := <- ch

	fmt.Println(x)
}
