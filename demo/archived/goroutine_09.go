package main

import (
	"fmt"
	"time"
)

func fibonacci(c, quit chan int) {
	//x, y := 0, 1
	x := 0
	for {
		select {
		case c <- x:
			//x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c, quit := make(chan int), make(chan int)

	go func(){
		for i:=0; i<10; i++ {
			fmt.Println("<-c num is ", i, "content is ", <-c, "time.Now(): ", time.Now())

		}
		quit <- 0
	}()
	fibonacci(c, quit)
}