package main

/*
	https://www.airs.com/blog/archives/277
*/

import (
	"fmt"
)

type S struct {
	i int
}

func (p *S) Get() int {
	return p.i
}

func (p *S) Set(v int) {
	p.i = v
}

type I interface {
	Get() int
	Set(int)
}

func f(p I) {
	fmt.Println(p.Get())
	p.Set(2)
	fmt.Println(p.Get())
}

func g(i interface{}) int {
	return i.(I).Get()
}

func h() {
	var s S
	fmt.Println(g(&s))
	//fmt.Println(g(s))
}

func main() {
	var s S
	f(&s)
	h()
}
