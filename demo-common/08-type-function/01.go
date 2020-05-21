package main

import "fmt"

// Greeting function types
// 函数类型是表示所有包含相同参数和返回类型的函数集合
// A function type denotes the set of all functions with the same parameter and result types
type Greeting func(name string) string

func english(name string) string {
	return "Hello, " + name
}

func say(g Greeting, n string) {
	fmt.Println(g(n))
}

func main() {
	say(english, "World")
}