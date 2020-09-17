package main

/*
	通过channel机制对每个协程都设置退出条件，从而达到回收资源的目的，其中channel是一个消息队列通道。
	https://yougg.github.io/2017/06/12/go%E8%AF%AD%E8%A8%80%E5%AE%89%E5%85%A8%E7%BC%96%E7%A8%8B%E8%A7%84%E8%8C%83/
*/

import (
	"fmt"
	"time"
)

func doWaiter(name string, second int, signal chan int) {
	for {
		select {
		case <-time.Tick(time.Duration(second) * time.Second):
			fmt.Println(name, " is ready!")
		case <-signal:
			fmt.Println(name, " close goroutine.")
			return
		}
	}
}


func main() {
	var signal1 = make(chan int)
	var signal2 = make(chan int)

	defer close(signal1)
	defer close(signal2)

	go doWaiter("Tea", 2, signal1)
	go doWaiter("Coffee", 1, signal2)

	fmt.Println("main() is waiting..." )
	time.Sleep(4 * time.Second)

	signal1 <- 1
	signal2 <- 1
	time.Sleep(time.Second)

}