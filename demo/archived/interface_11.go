package main

import (
	"fmt"
)

type Binary uint64

func main() {
	b := Binary(200)

	any := (interface{})(b)
	fmt.Println(any)
	fmt.Printf("%T ", any)
}