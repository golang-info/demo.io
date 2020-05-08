package main

import (
	"fmt"
)

/*
	https://blog.51cto.com/pmghong/2462678?source=dra
*/

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

func (sl *SingleLink) GetHead() *LinkNode{
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

	ptr := sl.GetHead()
	for ptr != nil {
		fmt.Println("Data: ", ptr.Data)
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
		sl.tail.Next = node
		node.Next = nil
		sl.tail = node
	}
	sl.size++
}

func (sl *SingleLink) InsertByIndex(index int, node *LinkNode) {
	if node == nil {
		return
	}

	if index == 0 {
		sl.InsertByHead(node)
	} else {
		if index > sl.Length() {
			return
		} else if index == sl.Length() {
			sl.InsertByTail(node)
		} else {
			preNode := sl.Search(index - 1)
			currentNode := sl.Search(index)
			preNode.Next = node
			node.Next = currentNode
			sl.size++
		}
	}
}

func (sl *SingleLink) DeleteByIndex(index int) {
	if sl.Length() == 0 || index > sl.Length() {
		return
	}

	if index == 0 {
		sl.head = sl.head.Next
	} else {
		preNode := sl.Search(index - 1)
		if index != sl.Length() - 1{
			nextNode := sl.Search(index).Next
			preNode.Next = nextNode
		} else {
			sl.tail = preNode
			preNode.Next = nil
		}
	}
	sl.size--
}

func (sl *SingleLink) DeleteByData(Data interface{}) {
	if sl.Length() == 0 || Data == nil {
		return
	}

	node := sl.head
	preNode := sl.head
	for node.Next != nil {
		preNode = node
		node = node.Next
		if node.Data.(int) == Data.(int) {
			preNode.Next = node.Next
			node.Next = nil
			node.Data = nil
			node = nil
			return
		}
	}
}

func (sl *SingleLink) Destroy(){
	sl.tail = nil
	sl.head = nil
	sl.size = 0
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

func main() {
	//初始化链表
	sl := InitSingleLink()

	//指定指标插入
	for i := 0; i < 5; i++ {
		snode := &LinkNode{
			Data: i,
		}
		sl.InsertByIndex(i, snode)
	}

	sl.Print()
	fmt.Println("========================================")

	var snode *LinkNode

	snode = &LinkNode{
		Data: 6,
	}
	sl.InsertByHead(snode)
	sl.Print()
	fmt.Println("==========================================")

	snode = &LinkNode{
		Data: 5,
	}
	sl.InsertByTail(snode)
	sl.Print()
	fmt.Println("===========================================")

	node := sl.Search(2)
	fmt.Println("Node2: ", node.Data)
	fmt.Println("============================================")

	sl.DeleteByIndex(2)
	sl.Print()
	fmt.Println("=============================================")

	sl.DeleteByData(3)
	sl.Print()
}