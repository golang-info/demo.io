package main

import "fmt"

func main() {
	var i int
	var p *int
	var pp **int
	i = 0
	p = &i
	pp = &p
	*p++
	fmt.Println(i, p, *p, pp, *pp)
}
