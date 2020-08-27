package main

import (
	"fmt"
	"sync"
)

func main() {
	ints := generate()
	results := make(chan int)
	go average(ints, results)
	for v := range results {
		fmt.Println("Average:", v)
	}
}

func generate() <-chan []int {
	out := make(chan []int)
	go func() {
		data := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		for _, v := range data {
			out <- v
		}
		close(out)
	}()

	return out
}

func average(i <-chan []int, receiver chan int) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for v := range i {
			receiver <- avg(v)
		}
		wg.Done()
	}()
	wg.Wait()
	close(receiver)
}

func avg(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	result := sum / len(nums)
	return result
}
