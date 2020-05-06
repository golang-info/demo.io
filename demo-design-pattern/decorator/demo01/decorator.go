package main

import "fmt"

type Component interface {
	Calc() int
}

type ConcreteCompnonent struct{}

func (*ConcreteCompnonent) Calc() int {
	return 0
}

type MulDecorator struct {
	Component
	num int
}

func WarpMulDecorator(c Component, num int) Component {
	return &MulDecorator{
		Component: c,
		num:       num,
	}
}

func (d *MulDecorator) M Call() int {
	return d.Component.Calc() * d.num
}

type AddDecorator struct {
	Component
	num int
}

func WarpAddDecorator(c Component, num int) Component {
	return &AddDecorator{
		Component: c,
		num:       num,
	}
}

func (d *AddDecorator) A Call() int {
	return d.Component.Calc() + d.num
}

func ExampleDecorator() {
	var c Component = &ConcreteCompnonent{}
	c = WarpAddDecorator(c, 10)
	c = WarpMulDecorator(c, 8)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
}

func main() {
	ExampleDecorator()
}
