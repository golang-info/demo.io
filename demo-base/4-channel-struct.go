package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func(ch chan struct{}){
		println("doing in go-func")
		time.Sleep(3 * time.Second)
		ch<- struct{}{}
	}(ch)

	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("time out.")
			return
		}
	}
	//<-ch
	//fmt.Printf("exit main, %T\n", <-ch)
}
