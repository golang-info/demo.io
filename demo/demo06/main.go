package main

import "fmt"

func jia(a, b int) int {
	c := a + b
	return c
}

func main() {
	d := jia(1, 2)
	fmt.Println(d)
}
