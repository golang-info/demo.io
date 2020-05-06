package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main(){
	c1 := boring("boring..., c1")
	c2 := boring("boring..., c3")
	for i := 0; i <5; i++ {
		fmt.Printf("You see: %q\n", <-c1)
		fmt.Printf("You see: %q\n", <-c2)
	}
	fmt.Println("You're boring. I'm leaving")
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}