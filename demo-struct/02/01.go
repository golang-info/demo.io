package main

import (
	"fmt"
	"reflect"
)

/*
	https://studygolang.com/articles/16507
*/

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age" valid:"1-100"`
}

func (p *Person) validation() bool {
	v := reflect.ValueOf(*p)
	tag := v.Type().Field(1).Tag.Get("valid")
	val := v.Field(1).Interface().(int)
	fmt.Printf("tag=%v, val=%v\n", tag, val)

	result := strings.Split(tag, "-")
	var min, max int
	min, _ = strconv.Atoi(result[0])
	max, _ = strconv.Atoi(result[1])

	if val >= min && val <= max {
		return true
	} else {
		return false
	}
}

func main() {
	person1 := Person{"tom", 12}
	if person1.validation() {
		fmt.Printf("person 1: valid\n")
	} else {
		fmt.Printf("person 1: invalid\n")
	}

	person2 := Person{"tom", 250}
	if person2.validation() {
		fmt.Printf("person 2: valid\n")
	} else {
		fmt.Printf("person 2: invalid\n")
	}
}
