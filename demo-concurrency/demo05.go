package main

import (
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		//wg.Add(1)

		go func(id int) {
			defer wg.Done()

			time.Sleep(time.Second)
			println("grounite", id, "done")
		}(i)
	}

	println("main...")
	wg.Wait()
	println("main exit.")
}
