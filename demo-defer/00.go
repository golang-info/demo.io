package main

import "fmt"

func main() {
	defer fmt.Println("I'm run after the function completes")
	fmt.Println("Hello World!")
}
