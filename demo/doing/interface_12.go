package main

// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

import (
	"fmt"
)

func PrintAll(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
} 

func main() {
	names := []string{"standey", "david", "oscar"}
	vals := make([]interface{}, len(names))
	for i, v := range names {
		vals[i] = v
	}
	PrintAll(vals)
}