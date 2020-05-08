package main

import "fmt"

type Vehicle interface {
	Structure()
	Speed()
}

type Hunman interface {
	Structure()
	Performance()
}

type Car string

func (c Car) Structure() []string {
	parts := []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}

type Man string 

func (m Man) Structure() []string {
	parts := []string{"Brain", "Heart", "Nose", "Eyelashes", "Stomach"}
	return parts
}

func main() {
	var c Car
	var m Man
	
	for i, j := range c.Structure() {
		fmt.Printf("%-15s <=====> %15s\n", j, m.Structure()[i])
	}
}