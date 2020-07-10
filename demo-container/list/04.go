package main

import (
	"container/list"
	"fmt"
)

func main() {
	var l *list.List
	var el *list.Element

	l = list.New()

	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)

	el = l.Front()

	fmt.Println("Going from the right\n")

	fmt.Printf("First element: %d\n", el.Value)
	fmt.Printf("Second element: %d\n", el.Next().Value)
	fmt.Printf("Third element: %d\n", el.Next().Next().Value)

	fmt.Println("Going form the left\n")

	el = l.Back()

	fmt.Printf("Firt element: %d\n", el.Value)
	fmt.Printf("Second element: %d\n", el.Prev().Value)
	fmt.Printf("Third element: %d\n", el.Prev().Prev().Value)

}
