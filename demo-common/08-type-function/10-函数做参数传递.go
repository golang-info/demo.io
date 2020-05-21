package main

import "fmt"

type CalculateType func(a, b int) int //声明了一个函数类型

// 加法函数
func add(a, b int) int {
	return a + b
}

// 乘法函数
func mul(a, b int) int {
	return a * b
}

func Calculate(a, b int, f CalculateType) int {
	return f(a, b)
}

func main() {
	a, b := 2, 3
	fmt.Println(Calculate(a, b, add))
	fmt.Println(Calculate(a, b, mul))
}
