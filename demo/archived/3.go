package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick", time.Now())
		case <-boom:
			fmt.Println("BOOM!", time.Now())
			return
		default:
			fmt.Println("     ")
			time.Sleep(40 * time.Millisecod)
		}
	}
}
