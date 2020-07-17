package main

import (
	"fmt"
	"sync"
)

type ArrayStack struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (stack *ArrayStack) Push(v string) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	stack.array = append(stack.array, v)
	stack.size = stack.size + 1
}

func (stack *ArrayStack) Pop() string {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	if stack.size == 0 {
		panic("empty")
	}

	v := stack.array[stack.size-1]

	//切片收缩，但可能占用空间越来越多
	//stack.array = stack.array[0:stack.size-1]

	//创建新的数组，空间占用不会越来越大，但可能移动元素次数过多
	newArray := make([]string, stack.size-1, stack.size-1)
	for i := 0; i < stack.size-1; i++ {
		newArray[i] = stack.array[i]
	}
	stack.array = newArray

	stack.size = stack.size - 1
	return v
}

func (stack *ArrayStack) Peek() string {
	if stack.size == 0 {
		panic("empty")
	}

	v := stack.array[stack.size-1]
	return v
}

func (stack *ArrayStack) Size() int {
	return stack.size
}

func (stack *ArrayStack) isEmpty() bool {
	return stack.size == 0
}

func main() {
	arrayStack := new(ArrayStack)
	arrayStack.Push("cat")
	arrayStack.Push("dog")
	arrayStack.Push("hen")
	fmt.Println("size", arrayStack.Size())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size", arrayStack.Size())
	arrayStack.Push("drag")
	fmt.Println("pop:", arrayStack.Pop())
	fmt.Println("size", arrayStack.Size())
}
