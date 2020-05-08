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
	return fmt.Sprintf("Moooohoooppp %s Boooo", c.name)
}

func (d *Dog) Call() string {
	return fmt.Sprintf("WooooofWaf %s Wuf!", d.name)
}

func main() {
	var animals []Animal = []Animal {
		&Cow{"Betsy"},
		&Dog{"Roover"},
		&Cow{"Bull"},
		&Dog{"Fiff"},
	}

	for _, animal := range animals {
		fmt.Println(animal.Call())
	}
}