package main

import (
	"encoding/json"
	"fmt"
)

type Foo struct {
	Bar string
	Baz string
}

func main() {
	f := Foo{
		"Joe Junior",
		"Hello Shabado",
	}

	b, _ := json.Marshal(f)

	fmt.Println(string(b))

	json.Unmarshal(b, &f)

	fmt.Println(b)
}
