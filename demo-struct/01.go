package main

/*
	golang 判断结构体为空
*/

import (
	"fmt"
	"reflect"
)

type A struct {
	name string
	age  int
}

func (a *A) IsEmpty() bool {
	return reflect.DeepEqual(a, A{})
}

func main() {
	var a A

	a = A{
		name: "ok",
		age:  14,
	}

	if a == (A{}) {
		fmt.Println("A == A{} empty")
	}

	if a != (A{}) {
		fmt.Println("A == A{} not empty")
	}

	if a.IsEmpty() {
		fmt.Println("reflect deep is empty")
	}

	if !a.IsEmpty() {
		fmt.Println("reflect deep is not empty")
	}
}
