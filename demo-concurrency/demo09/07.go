package main

import (
	"fmt"
	"sync"
)

func main() {
	ints := generate()
	c1 := average(ints)
	c2 := average(ints)

	for v := range merge(c1, c2) {
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

func average(in <-chan []int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- avg(v)
		}
		close(out)
	}()
	return out
}

func avg(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	result := sum / len(nums)
	return result
}

func merge(ch ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}

	wg.Add(len(ch))

	for _, c := range ch {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
