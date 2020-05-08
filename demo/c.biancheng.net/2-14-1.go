package main

import "fmt"

func main() {
	var cat int = 1
	var ptr string = "banana"
	fmt.Printf("%p\n%p\n", &cat, &ptr)
}
