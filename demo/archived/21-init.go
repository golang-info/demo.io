package main

import (
	"fmt"
	"unsafe"
)

func main() {
	i := int(31)
	fmt.Println(unsafe.Sizeof(i))
	fmt.Printf("%b\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%o\n", i)



	j := 1
	fmt.Println(unsafe.Sizeof(j))

	u := uint(1)
	fmt.Println(unsafe.Sizeof(u))
}
