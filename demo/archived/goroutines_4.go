package main

import (
	"fmt"
	"time"
)

/*
	very good demo!!!
*/

//This subroutine pushes numbers 0 to 9 to the channel and closes the channel
func add_to_channel(ch chan int) {
	fmt.Println("Send data")
	for i:=0; i<10; i++ {
		ch <- i  // pushing data to channel
	}
	close(ch) //closing  the channel
}

//This subroutine fetches data form the channel and prints it.
func fetch_from_channel(ch chan int) {
	fmt.Println("Read data")
	for {
		//fetch data from channel
		x, flag := <- ch
		//flag is true if data is received from the channel
		//flag is false when the channel is closed
		if flag == true {
			fmt.Println(x)
		} else {
			fmt.Println("Empty channel")
			break
		}
	}
}

func main() {
	//creating a channel variable to transport integer values
	ch := make(chan int)

	//invoking the subroutines to add and fetch form the channel
	//These routines execute simultaneously
	go add_to_channel(ch)
	go fetch_from_channel(ch)

	//delay is to prevent the exiting of main() before goroutines finish
	time.Sleep(5 * time.Second)
	fmt.Println("Inside main()")
}