package main

import "fmt"

func ExampleDecorator() {
	var c Component = &ConcreteCompnonent{}
	c = WarpAddDecorator(c, 10)
	fmt.Printf("res add: %d\n", c.Calc())
	c = WarpMulDecorator(c, 8)
	fmt.Printf("res mul: %d\n", c.Calc())

}

func main() {
	ExampleDecorator()
}
