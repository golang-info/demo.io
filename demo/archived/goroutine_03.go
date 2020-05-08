package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func(){
		fmt.Println("Go Go Go!!!")
		c <- true
		close(c)
	}()
	
	for v := range c {
		fmt.Println(v)
	}
}