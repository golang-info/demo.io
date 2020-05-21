package main

/*
	https://medium.com/@kandros/go-function-type-reusable-function-signatures-2389f6bdd4f6
*/

import "fmt"

type myFunctionType = func(a, b string) string

func main() {
	var explicit myFunctionType = func(a, b string) string {
		return fmt.Sprintf("%s %s", a, b)
	}

	implicit := func(a, b string) string {
		return fmt.Sprintf("%s %s", b, a)
	}

	functionTypeConsumer(explicit)
	functionTypeConsumer(implicit)
	functionTypeConsumer(func(a, b string) string {
		return fmt.Sprintf("%s %s", a, b)
	})
}

func functionTypeConsumer(fn myFunctionType) {
	s := fn("Hello", "World")
	fmt.Println(s)
}
