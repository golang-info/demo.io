package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l *list.List

	l = list.New()

	// root -> 1
	l.PushFront(1)

	// root -> 2 -> 1
	// 1 has been pushed by 2
	l.PushFront(2)

	// root -> 3 -> 2 -> 1
	l.PushFront(3)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	for e := l.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value)
	}
}
