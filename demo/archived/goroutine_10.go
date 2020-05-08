package main

import "fmt"

//两种channel的声明方式	
func main() {
	var ch01 chan int
	fmt.Println("value ch01: ", ch01)
	fmt.Printf("type ch01: %T\n", ch01)

	ch02 := make(chan int)
	fmt.Println("value ch01: ", ch02)
	fmt.Printf("type ch01: %T\n", ch02)
}