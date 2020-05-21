package main

// Go 语言的类型转换的基本格式：
// type_name(expression)

import "fmt"

type CalculateType func(int, int) //声明了一个函数类型

//该函数类型实现了一个方法
func (c *CalculateType) Serve() {
	fmt.Println("我是一个函数类型")
}

// 加法函数
func add(a, b int) {
	fmt.Println(a + b)
}

// 乘法函数
func mul(a, b int) {
	fmt.Println(a * b)
}

func main() {
	a := CalculateType(add) // 将add 函数强制转换成CalculateType类型
	b := CalculateType(mul) // 将mul 函数强制转换成CalculateType类型
	a(2, 3)
	b(2, 3)
	a.Serve()
	b.Serve()
}
