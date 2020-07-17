package main

import (
	"fmt"
	"sync"
)

type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

func (stack *LinkStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.root == nil {
		stack.root = new(LinkNode)
		stack.root.Value = v
	} else {
		preNode := stack.root

		newNode := new(LinkNode)
		newNode.Value = v

		newNode.Next = preNode

		stack.root = newNode
	}

	stack.size = stack.size + 1
}

func (stack *LinkStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty")
	}

	topNode := stack.root
	v := topNode.Value

	stack.root = topNode.Next

	stack.size = stack.size - 1

	return v
}

func (stack *LinkStack) Peek() string {
	if stack.size == 0 {
		panic("empty")
	}

	v := stack.root.Value
	return v
}

func (stack *LinkStack) Size() int {
	return stack.size
}

func (stack *LinkStack) IsEmpty() bool {
	return stack.size == 0
}

func main() {
	LinkStack := new(LinkStack)
	LinkStack.Push("cat")
	LinkStack.Push("dog")
	LinkStack.Push("hen")
	fmt.Println("size", LinkStack.Size())
	fmt.Println("pop:", LinkStack.Pop())
	fmt.Println("pop:", LinkStack.Pop())
	fmt.Println("size", LinkStack.Size())
	LinkStack.Push("drag")
	fmt.Println("pop:", LinkStack.Pop())
}
