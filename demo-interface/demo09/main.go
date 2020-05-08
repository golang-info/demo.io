package main

/*
	https://marcofranssen.nl/go-interfaces-and-type-assertions/
*/
import (
	"fmt"
)

func main() {
	var i interface{} = "Howdy fellow Gophers"

	//Type asserting test string
	//Howdy fellow Gophers true
	fmt.Println("Type asserting test string")
	s, ok := i.(string)
	fmt.Println(s, ok)
	fmt.Println()

	//Type asserting test float64
	//0 false
	fmt.Println("Type asserting test float64")
	f, ok := i.(float64)
	fmt.Println(f, ok)
	fmt.Println()

	//Type asserting byte slice
	//[] false
	fmt.Println("Type asserting byte slice")
	b, ok := i.([]byte)
	fmt.Println(b, ok)
}
