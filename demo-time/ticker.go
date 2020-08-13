package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(2 * time.Second)
	for t := range c {
		fmt.Printf("The time is now %v\n", t)
	}
}
