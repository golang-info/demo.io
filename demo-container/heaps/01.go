package main

import "fmt"

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swwap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func main() {
	fmt.Println("not to be used yet: :)")
}
