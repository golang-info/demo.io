package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)

	go func() {
		time.Sleep(3 * time.Second)
		timeout <- true
	}()

	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i
		}
	}()

	for {
		select {
		case m := <-ch:
			fmt.Println("ch: ", m)
		case t := <-timeout:
			fmt.Println("timeout: ", t)
			return
			case <-time.After(1 * time.Second):
				fmt.Println("timeout 00")
		}
	}

}
