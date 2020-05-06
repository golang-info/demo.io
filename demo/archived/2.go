package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)

	ch <- "hello"
	ch <- "world"
	//fatal error: all goroutines are asleep - deadlock!
	//ch <- "end"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)

}
