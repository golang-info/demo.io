package main

import (
	"fmt"
	"time"
)

func Ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}

func main() {
	go Ready("Tea", 2)
	go Ready("Coffee", 1)
	fmt.Println("I'm waitting.")
	time.Sleep(5 * time.Second)
}
