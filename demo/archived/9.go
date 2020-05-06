package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func funIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func(){ for { c <- <-input1}}()
	go func(){ for { c <- <-input2}}()
	return c
}

func main() {
	c := funIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring, i am leaving")
}