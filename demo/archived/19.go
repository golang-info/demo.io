package main

import "time"

func main() {
	timeout := make(chan bool)
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch <- 1
	}()

	for {
		select {
		case value := <-ch:
			println("ch: ", value)
			continue
			case <-timeout:
				println("timeout: 2")
				return
		}
	}
}
