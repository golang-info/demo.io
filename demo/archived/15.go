package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go power1()

	for {
		time.Sleep(time.Duration(1) * time.Minute)
	}
}

func power1() {
	var buf [1024]byte
	fmt.Println("power1.....")

	n := runtime.Stack(buf[:], true)

	fmt.Println(string(buf[:]), n)
}
