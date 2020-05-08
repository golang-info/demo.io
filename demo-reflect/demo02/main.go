package main

import (
	"reflect"
)

type INT int

type A struct {
	a int
}

type B struct {
	b string
}

type Ita interface {
	String() string
}

func main() {
	var a INT = 12
	var b int = 14

	ta := reflect.TypeOf(a)
	tb := reflect.TypeOf(b)

	if ta == tb {
		println("ta==tb")
	} else {
		println("ta!=tb")
	}

/*	c := int(a)
	println("c: ", c)
	println(reflect.TypeOf(c).Name())*/

	println(ta.Name())
	println(tb.Name())

	println(ta.Kind().String())
	println(tb.Kind().String())

	s1 := A{1}
	s2 := B{"tata"}

	println(reflect.TypeOf(s1).Name())
	println(reflect.TypeOf(s2).Name())

	println(reflect.TypeOf(s1).Kind().String())
	println(reflect.TypeOf(s2).Kind().String())

	ita := new(Ita)
	//var itb Ita = s2

	println(reflect.TypeOf(ita).Elem().Name())
	println(reflect.TypeOf(ita).Elem().Kind().String())

	//println(reflect.TypeOf(itb).Name())
	//println(reflect.TypeOf(itb).Kind().String())
}
