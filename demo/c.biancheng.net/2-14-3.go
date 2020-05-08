package main

import "fmt"

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	x, y := 1, 2

	swap(&x, &y)

	fmt.Println(x, y)
}
