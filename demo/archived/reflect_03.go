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

func main() {
	u := User{1, "OK", 12}
	Set(&u)
	fmt.Println(u)
}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("can not set")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")

	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	if  f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}