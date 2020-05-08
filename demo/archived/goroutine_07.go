package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("start...\n")
	go DoSomething(&wg)
	fmt.Println("end...\n")

	wg.Wait()

}

func DoSomething(wg *sync.WaitGroup) {
	fmt.Println("do something.")

	wg.Done()

}