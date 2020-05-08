package main

import "fmt"

type Geometry interface {
	Edges() int
}

type Polygons interface {
	Geometry
}

type Pentagon struct{}
type Hexagon struct{}
type Octagon struct{}
type Decagon struct{}

func (p Pentagon) Edges() int {return 5}
func (h Hexagon) Edges() int {return 6}
func (o Octagon) Edges() int {return 8}
func (d Decagon) Edges() int {return 10}

func main() {
	p := new(Pentagon)
	h := new(Hexagon)
	o := new(Octagon)
	d := new(Decagon)

	polygons := [...]Polygons{p, h, o, d}

	for i := range polygons {
		fmt.Println(polygons[i].Edges())
	}
}