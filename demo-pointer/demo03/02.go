package main

// 以指针引用的方式复制结构体
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
		Name: "Lemonade",
		Ice:  true,
	}

	b := &a
	b.Ice = false
	fmt.Printf("%+v\n", *b)
	fmt.Printf("%+v\n", a)
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a)
}
