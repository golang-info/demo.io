package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l *list.List
	var nums []int
	l = list.New()
	nums = []int{1, 2, 3, 4, 5}
	for _, n := range nums {
		l.PushFront(n)
	}

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
