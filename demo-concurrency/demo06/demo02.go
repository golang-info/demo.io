package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	nums := make([]int, 100)

	go func() {
		time.Sleep(time.Second)
		for i := 0; i < len(nums); i++ {
			nums[i] = i
		}
		ch <- struct{}{}
	}() 

	<-ch
	fmt.Println(nums)
}