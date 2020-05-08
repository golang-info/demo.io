package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id int
	Name string
	Age int
}

func (u User) Hello(name string) {
	fmt.Println("Hello " + name + ",", " my name is", u.Name)
}

func main() {
	u := User{1, "OK", 12}
	//u.Hello("Joe")

	v := reflect.ValueOf(u)
	mv := v.MethodByName("Hello")

	args := []reflect.Value{reflect.ValueOf("John")}
	mv.Call(args)
}