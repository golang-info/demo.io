package main

import "fmt"

type Polygons interface {
	Perimeter()
}

type Object interface {
	NumberOfSide()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println(5 * p)
}

func (p Pentagon) NumberOfSide() {
	fmt.Println("Pentagon has 5 sides.")
}

func main() {
	obj := Pentagon(50)
	obj.Perimeter()
	obj.NumberOfSide()
}