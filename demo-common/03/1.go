package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)

	runtime.GOMAXPROCS(100)

	for i:= 0; i < 100; i++ {
		go saycall(i, &wg)
	}

	wg.Wait()
}

func saycall(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello", i)
}


