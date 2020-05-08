package main

//类型系统 88page

import "fmt"

type T struct {
	a int
}

func (t T) Get() int {
	return t.a
}

func (t *T) Set(i int) {
	t.a = i
}

func (t *T) Print() {
	fmt.Printf("%p, %v, %d \n", t, t, t.a)
}

func main() {
	var t = &T{}

	//method value
	f := t.Set

	// 方法值调用
	f(2)
	t.Print()

	//方法值调用
	f(3)
	t.Print()
}
