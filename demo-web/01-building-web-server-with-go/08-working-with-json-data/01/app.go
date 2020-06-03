package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Address struct {
	Street string
	Landmark string
	// Pincode int  // encoding/json package only accesses the exported fields of struct types
	pincode int
}

func main() {
	a := Address{
		Street: "Viman Nagar",
		Landmark: "Nexa",
		// Pincode: 411014,
		pincode: 411014,
	}

	b := Address{}

	bytes, err := json.Marshal(a)
	if err != nil {
		log.Fatal("unable to encode")
	}
	fmt.Println(bytes)

	err = json.Unmarshal(bytes, &b)
	if err != nil {
		log.Fatal("unable to decode")
	}
	fmt.Println(b)
}


