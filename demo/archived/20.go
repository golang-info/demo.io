package main

import "fmt"

func main() {
	var set map[string]struct{}

	//Initialize the set
	set = make(map[string]struct{})

	//Add some values to the set:
	set["red"] = struct{}{}
	set["blue"] = struct{}{}

	//Check if a value is in the map:
	value, ok := set["red"]
	fmt.Print("Is red in the map?\n", ok)
	fmt.Println("value: ", value)

	_, ok = set["green"]
	fmt.Println("\nIs green in the map?\n", ok)
}


