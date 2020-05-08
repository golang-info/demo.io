package main

import (
	"fmt"
)

var pr = fmt.Println
var f = func() {
	pr("hello")
}

var g func()
func main() {
	g = f
	g()
}