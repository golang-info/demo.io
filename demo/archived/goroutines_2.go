package main

import (
	"fmt"
	"time"
)

func display() {
	time.Sleep(5 * time.Second)
	fmt.Println("Inside display()")
}

func main() {
	go display()
	fmt.Println("Inside main()")
}