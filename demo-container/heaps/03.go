package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i] , h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0: n - 1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	fmt.Println(h)

	// sorted when heap.Init()
	heap.Init(h)
	fmt.Println(h)

	fmt.Println((*h)[0])
	fmt.Println((*h)[1])
	fmt.Println((*h)[2])

	// the entire heap is sorted when you call heap.Push()
	heap.Push(h, 10)
	fmt.Println("Pushing number 10 to the heap")

	fmt.Println((*h)[0])
	fmt.Println((*h)[1])
	fmt.Println((*h)[2])
	fmt.Println((*h)[3])

	fmt.Println("Popping the first item again sorts the heap")

	heap.Pop(h)

	fmt.Println(h)

	fmt.Println((*h)[0])
	fmt.Println((*h)[1])
	fmt.Println((*h)[2])




}
