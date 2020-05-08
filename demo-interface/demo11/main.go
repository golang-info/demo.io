package main

import (
	"fmt"
)

type Animal interface {
	Call() string
}

type Cow struct {
	name string
}

type Dog struct {
	name string
}

func (c *Cow) Call() string {
	return fmt.Sprintf("MoooooBooooo %s Boool", c.name)
}

func (d *Dog) Call() string {
	return fmt.Sprintf("WooooWalf %s Walf!", d.name)
}

func main() {
	var animals []Animal = []Animal{
		&Cow{"Betsy"},
		&Dog{"Roover"},
		&Cow{"Bull"},
		&Dog{"Fifi"},
	}

	for _, animal := range animals {
		fmt.Println(getSpecies(animal))
		fmt.Println(animal.Call())
		fmt.Println()
	}
}

func getSpecies(i interface{}) string {
	var m string
	switch a := i.(type) {
	case Cow:
	case *Cow:
		m = fmt.Sprintf("I'm a cow, %#v", a)
		break
	case Dog:
	case *Dog:
		m = fmt.Sprintf("I'm a dog, %#v", a)
		break
	default:
		m = "Don't know what species this is"
	}
	return m
}
