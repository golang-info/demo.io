package main

import (
	"fmt"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(4)
}

func main() {
	exit := make(chan int)
	go func() {
		defer  close(exit)
		go func() {
			fmt.Println("b")
		}()
	}()

	for i := 0; i < 4; i++ {
		fmt.Println("a:", i)

		if i == 1 {
			runtime.Gosched() //切换任务
		}
	}
	<-exit
}
