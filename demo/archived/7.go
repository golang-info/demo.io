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
func boring(msg string, c  chan string) {
	for i := 0; ; i++ {
		//fmt.Println(msg, i)
		//time.Sleep(time.Second)
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func main() {
	c := make(chan string)
	go boring("Boring...", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)  //Receive expression is just value.
	}
	fmt.Println("I am listening")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring, I'm leaving.")
}
