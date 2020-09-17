package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	var any interface{}

	any = true
	output(any)

	any = 12.14
	output(any)

	any = "hello"
	output(any)

	any = []int{1, 2, 4}
	output(any)

	any = [5]int{1, 2, 3, 4, 5}
	output(any)

	any = map[string]int{"hello": 1}
	output(any)

	any = new(bytes.Buffer)
	output(any)
}


func output(i interface{}) {
	fmt.Println(i, " --> ", reflect.TypeOf(i))

}

