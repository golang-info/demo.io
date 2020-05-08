package main

import "fmt"

/*
	A.fun() -> B.f()
*/

type iA interface {
	fun()
}

type A struct {
	//...
	b B
}

func (a A) fun() {
	//...
	fmt.Println("A.fun()")
	a.b.f()
}


type B struct {
	//...
}

func (b B) f() {
	fmt.Println("B.f()")
}

func main() {
	var i iA = A{}
	i.fun()

}