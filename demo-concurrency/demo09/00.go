package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		c <- 43
	}()

	fmt.Println("Value of channal c:", <-c)
}
