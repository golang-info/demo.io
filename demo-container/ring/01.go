package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var r *ring.Ring

	// Create 5 ring.Ring(5)
	r = ring.New(5)

	// the inital value of the first ring is <nil>
	fmt.Println(r.Value)
	fmt.Println(r.Next().Value)

	// we call repeat this 5 times since I created a ring with 5 rings in it
}
