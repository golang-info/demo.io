package main

import "fmt"

/* 
	https://www.golangprograms.com/go-language/interface.html
*/

type Polygons interface {
	Perimeter()
}

type Pentagon int

func (p Pentagon) Perimeter() {
	fmt.Println(5 * p)
}

type Hexagon int

func (h Hexagon) Perimeter() {
	fmt.Println(6 * h)
}

func main() {
	var object Polygons

	object = Pentagon(50)
	object.Perimeter()

	object = Hexagon(50)
	object.Perimeter()
}