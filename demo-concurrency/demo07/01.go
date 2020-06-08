package main

import "fmt"

func main() {
	go printSomething()
	fmt.Println("1")
}

func  printSomething() {
	fmt.Println("0")
}