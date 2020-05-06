package main

import "fmt"
import "time"

func main() {
	go func() {
		fmt.Println("ok")
	}()
	time.Sleep(time.Second)
}
