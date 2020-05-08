package main

import (
	"fmt"
)

func main() {
	i := 7
	//inc(i)
	inc(&i)
	fmt.Println(i)
}

/* func inc(x int) {
	x++
} */

func inc(x *int) {
	*x++
}