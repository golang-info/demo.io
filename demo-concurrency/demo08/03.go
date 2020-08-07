package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	counter := 0

	const gs = 5
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()

			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println("count: ", counter)

			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
