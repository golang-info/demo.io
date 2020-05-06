package main

import (
	"fmt"
	"runtime"
)

func main() {
	buf := make([]byte, 1 << 20)
	runtime.Stack(buf, true)
	fmt.Printf("\n%s", buf)

	fmt.Println("hello")
}
