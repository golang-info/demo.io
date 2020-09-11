package main

import (
	"fmt"
)

type Speaker interface {
	Speak(s string)
}

type Human struct {
	Name string
}

func (xh Human) Speak(s string) {
	fmt.Printf("My name is %s, I can say %s\n", xh.Name, s)

}

type Animal struct {
	Name string
}

func (Am Animal) Speak(s string) {
	fmt.Printf("My name is %s, I can say %s\n", Am.Name, s)

}

func Do(s Speaker, h string) {
	s.Speak(h)
}

func main() {
	xh := Human{
		Name: "XiaoHong",
	}

	Am := Animal{
		Name: "xiaogou",
	}

	Do(xh, "hahahah")
	Do(Am, "wangwangwang")
}
