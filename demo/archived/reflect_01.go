package main

//无闻： Go编程基础

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello() {
	fmt.Println("Hello World.")
}

func main() {
	u := User{1, "OK", 12}
	info(u)
}

func info(o interface{}){
	t := reflect.TypeOf(o)
	fmt.Println("Type: ", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields: ")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
