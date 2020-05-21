package main

import "fmt"

const (
	FlagNone = 1 << iota
	FlagRed
	FlagGreen
	FlagBlue
)

func main() {
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
	fmt.Printf("%b %b %b\n", FlagRed, FlagGreen, FlagBlue)
}
