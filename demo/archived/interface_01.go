package main

import (
	"fmt"
)

type shape interface {
	area() float32
	perimeter() float32
}

type rectange struct {
	width float32
	height float32
}

func (rect rectange) area() float32 {
	return rect.width * rect.height
}

func (rect rectange) perimeter() float32 {
	return 2 * rect.width + 2 * rect.height
}

func display(shape shape) {
	fmt.Println(shape)
	fmt.Println(shape.area())
	fmt.Println(shape.perimeter())
}

func main() {
	rect := rectange{
		width: 2, 
		height: 3,
	}
	display(rect)
}

