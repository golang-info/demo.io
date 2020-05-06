package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Declaring and initializing
/*var c chan int
c = make(chan int)

c := make(chan int)*/

//Sending on a channel
//c <- 1

// Receiving from a channel
// The "arrow" indicates the direction of data flow
//value <-c
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		//time.Sleep(time.Second)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	go boring("Boring...")
	fmt.Println("I am listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring, I'm leaving.")
}
