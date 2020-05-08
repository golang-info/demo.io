package main

import (
	"fmt"
)

func main() {
	go Go()
}

func Go() {
	fmt.Println("Go Go Go!!!!")
}