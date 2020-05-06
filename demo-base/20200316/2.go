package main

import "fmt"

type Node struct {
	Value int
	Next *Node
}

func IsEmpty(list *Node) bool {
	return list == nil
}

func IsLast(position *Node, list *Node) bool {
	return position.Next == nil
}

func Find(value int, list *Node) *Node {
	p := list
	for p.Value != value {
		p = p.Next
	}
	return p
}

func Insert(value int, list *Node, position *Node) {
	tempCell := new(Node)
	if tempCell == nil {
		fmt.Println("out of space!")
	}
	tempCell.Value = value
	tempCell.Next = position.Next
	position.Next = tempCell
}

func Delete(value int, list *Node) {
	pPreV := FindPreVious(value, list)
	p := Find(value, list)
	pPreV.Next = p.Next
	p = nil
}

func FindPreVious(value int, list *Node) *Node {
	p := list
	for p.Next != nil && p.Next.Value != value {
		p = p.Next
	}
	return p
}

func PrintList(list *Node) {
	p := list
	for p != nil {
		fmt.Printf("%d-%p-%p", p.Value, p, p.Next)
	}
	fmt.Println()
}

func DeleteList(list **Node) {
	p := list
	for *p != nil {
		tmp := *p
		*p = nil
		*p = tmp.Next
	}
}

func main() {
	headNode := &Node{
		Value: 0,
		Next: nil,
	}

	list := headNode
	//Insert(1, list, headNode)
	PrintList(list)
}
