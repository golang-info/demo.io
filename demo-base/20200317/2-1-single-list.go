package main

import "fmt"

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

type SingleLink struct {
	head *LinkNode
	tail *LinkNode
	size int
}

func InitSingleLink() (*SingleLink) {
	return &SingleLink{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (sl *SingleLink) GetHead() *LinkNode {
	return sl.head
}

func (sl *SingleLink) GetTail() *LinkNode {
	return sl.tail
}

func (sl *SingleLink) Length() int {
	return sl.size
}

func (sl *SingleLink) Print() {
	fmt.Println("SingleLink size: ", sl.Length())
	if sl.size == 0 {
		return
	}

	fmt.Println("sl.head: ", sl.GetHead())
	ptr := sl.GetHead()
	for ptr != nil {
		if ptr.Next != nil {
			fmt.Println("Data: ", ptr.Data, "ptr: ", ptr, "ptr.Next: ", ptr.Next, "*ptr.Next: ", *ptr.Next, "sl: ", sl)
		} else {
			fmt.Println("Data: ", ptr.Data, "ptr: ", ptr, "ptr.Next: ", ptr.Next, "*ptr.Next: ", "0 nil")
		}
		ptr = ptr.Next
	}
}

func (sl *SingleLink) InsertByHead(node *LinkNode) {
	if node == nil {
		return
	}

	if sl.Length() == 0 {
		sl.head = node
		sl.tail = node
		node.Next = nil
	} else {
		oldHeadNode := sl.GetHead()
		sl.head = node
		sl.head.Next = oldHeadNode
	}
	sl.size++
}

func (sl *SingleLink) InsertByTail(node *LinkNode) {
	if node == nil {
		return
	}

	if sl.size == 0 {
		sl.head = node
		sl.tail = node
		node.Next = nil
	} else {
		oldHeadNode := sl.GetHead()
		sl.head = node
		sl.head.Next = oldHeadNode
	}
	sl.size++
}

func (sl *SingleLink) Search(index int) (node *LinkNode) {
	if sl.Length() == 0 || index > sl.Length() {
		return nil
	}

	if index == 0 {
		return sl.GetHead()
	}

	node = sl.head
	for i := 0; i <= index; i++ {
		node = node.Next
	}
	return
}

func (sl *SingleLink) InsertByIndex(index int, node *LinkNode) {
	if node == nil {
		return
	}

	if index == 0 {
		sl.InsertByHead(node)
	} else if index == sl.Length() {
		sl.InsertByTail(node)
	} else {
		proNode := sl.Search(index - 1)
		currentNode := sl.Search(index)
		proNode.Next = node
		node.Next = currentNode
		sl.size++
	}
}

func main() {
	sl := InitSingleLink()

	for i := 0; i < 5; i++ {
		snode := &LinkNode{
			Data: i,
		}
		sl.InsertByIndex(i, snode)
	}

	sl.Print()
	fmt.Println("--------------------------------------")
}


