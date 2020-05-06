package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(s)
	fmt.Println(s[1:2])
	fmt.Println(s[:])
	fmt.Println(s[:1])
	fmt.Printf("%T\n", s)
	fmt.Println(len(s), " ", cap(s))

	t := [3]int{0, 1, 2}
	fmt.Printf("%T\n ", t)
	fmt.Println(len(t))

	q := make([]int, 2, 4)
	fmt.Println(q, len(q), cap(q))

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) //2**i
	}
	fmt.Println(pow)

	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
