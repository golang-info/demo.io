package main

import (
	"fmt"
	"time"
)

/*
	awsome!!!
*/

//push data to channel with a 4 second delay
func data1(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "from data1()"
}

//push data to channel with a 2 second delay
func data2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "form data2()"
}

func main() {
	//creating channel variables for transporting string values
	chan1 := make(chan string)
	chan2 := make(chan string)

	//invoking the subroutines with channel variables
	go data1(chan1)
	go data2(chan2)

	//Both case statements wait for data in the chan1 or chan2
	//chan2 gets data first since the delay is only 2 sec in data2()
	//So the second case will execute and exits the select block
	select {
	case x := <-chan1:
		fmt.Println(x)
	case y := <-chan2:
		fmt.Println(y)
	/*
		Both case statements check for data in chan1 or chan2
		But data is not available (both routines have a delay of 2 and 4 sec)
		So the default block will be executed without waiting for data in channels.
	*/
	//default:
	//	fmt.Println("Default case executed")
	}
}