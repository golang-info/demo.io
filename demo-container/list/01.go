package main

import (
	"container/list"
	"fmt"
)

func main() {
	// first way is with the provided New function
	l := list.New()
	fmt.Printf("Instance of an empty list %v\n", l)
}
