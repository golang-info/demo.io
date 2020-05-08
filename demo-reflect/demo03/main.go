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

func (this User) String() {
	println("User:", this.Id, this.Name, this.Age)
}

func Info(o interface{}) {
	v := reflect.ValueOf(o)

	t := v.Type()

	println("Type:", t.Name())

	println("Fields:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		switch value := value.(type) {
		case int:
			fmt.Printf(" %6s: %v = %d\n", field.Name, field.Type, value)
		case string:
			fmt.Printf(" %6s: %v = %s\n", field.Name, field.Type, value)
		default:
			fmt.Printf(" %6s: %v = %s\n", field.Name, field.Type, value)
		}
	}
}

func main() {
	u := User{1, "Tom", 30}
	Info(u)
}
