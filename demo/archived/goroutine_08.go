package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("worker %d starting at %d ...\n", id, time.Now())
	time.Sleep(time.Second)
	fmt.Printf("worker %d done at %d ...\n", id, time.Now())
} 

func main() {
	var wg sync.WaitGroup

	for i:=1; i<=5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}