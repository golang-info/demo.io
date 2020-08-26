package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				fmt.Println("Result:", square(n))
			}
		}
	}()

	time.Sleep(time.Second * 3)

	fmt.Println("cancelling context...")
	cancel()
	fmt.Println("context cancelled!")

	time.Sleep(time.Second * 3)
}

func square(n int) int {
	time.Sleep(time.Millisecond * 200)
	return n * n
}
