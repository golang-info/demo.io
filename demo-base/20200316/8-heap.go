package main

import (
	"container/heap"
	"fmt"
)

type IntergerHeap []int

func (iheap IntergerHeap) Len() int { return len(iheap) }

func (iheap IntergerHeap) Less(i, j int) bool { return iheap[i] < iheap[j] }

func (iheap IntergerHeap) Swap(i, j int) { iheap[i], iheap[j] = iheap[j], iheap[i] }

func (iheap *IntergerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

func (iheap *IntergerHeap) Pop() interface{} {
	var n int
	var x1 int
	var provious IntergerHeap = *iheap
	n = len(provious)
	x1 = provious[n-1]
	*iheap = provious[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntergerHeap = &IntergerHeap{1, 4, 5}

	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}