package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func(){
		fmt.Println("Go Go Go!!!")
		c <- true
	}()
	//time.Sleep(1 * time.Second)
	<- c
}