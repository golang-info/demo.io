package main

import "fmt"

type emp struct {
	name string
	address string
	age int
}

func(e emp) display() {
	fmt.Println(e.name)
	fmt.Println(e.address)
	fmt.Println(e.age)
}

func main() {
	var empdata1 emp

	empdata1.name = "John"
	empdata1.address = "Street-1. Lodon"
	empdata1.age = 30

	empdata2 := emp {
		"Raj", "Building-1, Paris", 25,
	}

	empdata1.display()
	empdata2.display()
}