package main

import (
	"fmt"
)

type Animal interface {
	Speaking() string
}

type Dog struct {

}

func (d *Dog) Speaking() string {
	return "Woof"
}

type Cat struct {

}

func (c Cat) Speaking() string {
	return "Meow"
}

type Llama struct {

}

func (l Llama) Speaking() string {
	return "?????"
}

type Javaprogramer struct {

}

func (j Javaprogramer) Speaking() string {
	return "Design patterns!"
}

func main() {
	animals := []Animal{&Dog{}, &Cat{}, new(Llama), Javaprogramer{}}
	for _, animal := range animals {
		fmt.Println(animal.Speaking())
	}
}