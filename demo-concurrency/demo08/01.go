package main

import "fmt"

func main() {
	chat := make(chan string, 5)

	chat <- "a"
	chat <- "b"
	chat <- "c"
	chat <- "d"
	chat <- "e"
	//chat <- "f"

	fmt.Println(<-chat)
	fmt.Println(<-chat)
	fmt.Println(<-chat)
	fmt.Println(<-chat)
	fmt.Println(<-chat)
	//fmt.Println(<-chat) // fatal error: all goroutines are asleep - deadlock!
}
