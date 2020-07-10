package main

import (
	"container/ring"
	"fmt"
)

func main() {
	// Create a new ring of size 5
	r := ring.New(5)

	n := r.Len()

	// Initialize the ring with some integer values
	for i := 0; i <n; i++ {
		r.Value = i
		r = r.Next()
	}

	// fmt.Println(r.Value)
	// fmt.Println(r.Next().Value)
	// fmt.Println(r.Next().Next().Value)

	// This is copy/paste straight from the source code.
	// Notice p != r. r is the start ring and p is the next
	// When p becomes r, it stops
	// for p := r.Next(); p != r; p = p.Next() {
	//  fmt.Println(p.Value)
	// }

	r.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
	
	s := r.Next().Next().Next()
	s.Do(func(p interface{}) {
		fmt.Println(p.(int))
	})
}
