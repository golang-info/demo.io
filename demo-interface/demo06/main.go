package main

import (
	"fmt"
	"reflect"
)

func main() {
	var anyType interface{}
	anyType = "Canada"
	fmt.Println("Variable type: ", reflect.TypeOf(anyType))

	str, ok := anyType.(string)
	if ok {
		fmt.Println("Variable type: ", reflect.TypeOf(str))
	}	else {
		fmt.Println("Variable is not String.")
	}

	var intType = 100
	anyType = intType
	fmt.Println("Variable type: ", reflect.TypeOf(anyType))

	var anotherType interface{}
	anotherType = anyType
	fmt.Println("Variable type: ", reflect.TypeOf(anotherType))
}

