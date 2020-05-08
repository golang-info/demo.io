package main

import "fmt"

type Data struct {
}

func dummy() *Data {
	var c Data
	return &c
}

func main() {
	fmt.Println(dummy())
}
