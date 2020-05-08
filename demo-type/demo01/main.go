package main

//类型系统 87page

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

func main() {

	var t = &T{}

	t.Set(2)
	b := t.Get()
	fmt.Println(b)

}
