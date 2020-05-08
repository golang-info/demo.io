package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan struct{}, 2)


	var l sync.Mutex
	go func() {
		l.Lock()
		defer l.Unlock()
		fmt.Println(time.Now())
		fmt.Println("goroutine1: 我会锁定大概2s")
		time.Sleep(time.Second * 2)
		fmt.Println("goroutine1: 我解锁了， 你们去抢吧")
		ch <- struct{}{}
	}()

	go func() {
		fmt.Println(time.Now())
		fmt.Println("gorountine2: 等待解锁")
		l.Lock()
		defer l.Unlock()
		fmt.Println(time.Now())
		fmt.Println("gorountine2: 哈哈， 我锁定了")

		ch <- struct{}{}
	}()

	// 等待 gorountine 执行结束
	for i := 0; i < 2; i++ {
		<-ch
	}
}