package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		//计数加1
		wg.Add(1)
		go func(i int){
			//计数减1
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(i))
			fmt.Printf("goroutine%d 结束\n", i)
		}(i)
	}

	//等待结束
	wg.Wait()
	fmt.Println("所有 gorountine 执行结束")
}