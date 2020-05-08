package main

import "fmt"

func main() {
	ch := make(chan string, 5)
	ch <- "h1"
	ch <- "h2"
	ch <- "h3"
	ch <- "h4"

	fmt.Println("Lengh of the channel is ", len(ch))
	fmt.Println("Capacity of the channel is ", cap(ch))
}