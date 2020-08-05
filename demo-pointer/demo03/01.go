package main

// 以值应用的方式复制结构体
// Go 语言入门经典

import (
	"fmt"
)

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lenonade",
		Ice:  true,
	}
	b := a
	b.Ice = false

	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p \n", &a)
	fmt.Printf("%p \n", &b)
}
