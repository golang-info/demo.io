package main

import "fmt"

type Visitor func(name string) (shouldContinue bool)

var ok string

func main() {
	ok := "hhhh"
	if !Visitor(ok) {
		fmt.Println("not ok")
	} else {
		fmt.Println("ok")
	}
}
