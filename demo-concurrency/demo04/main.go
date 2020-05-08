package main

import (
	"fmt"
	"time"
)

func main() {
	go beeg()
	fmt.Println("Listenning for the beeps")
	time.Sleep(2 * time.Second)
	fmt.Println("That Beep is anoying , quit listenning!")
}

func beeg() {
	for i := 0;; i++ {
		fmt.Println("beep", i)
		time.Sleep(time.Second)
	}
}
