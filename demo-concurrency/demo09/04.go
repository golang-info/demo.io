package main

import (
	"fmt"
)

func main() {
	ints := generate()
	results := average(ints)
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

func average(i <-chan []int) <-chan int {
	out := make(chan int)
	go func() {
		for nums := range i {
			out <- avg(nums)
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
