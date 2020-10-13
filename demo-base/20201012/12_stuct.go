package main

import (
	"fmt"
	"unsafe"
)

type Set map[string]struct{}

func main() {
	fmt.Println(unsafe.Sizeof(struct{}{}))

	set := make(Set)

	for _, item := range []string{"A", "A", "B", "C"} {
		set[item] = struct{}{}
	}

	fmt.Println(len(set))
	if _, ok := set["A"]; ok {
		fmt.Println("A exists")
	}
}
