package main

import "sync"
import "fmt"

type LinkQueue struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

func (queue *LinkQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		newNode := new(LinkNode)
		newNode.Value = v

		nowNode := queue.root
		for nowNode.Next != nil {
			newNode = newNode.Next
		}

		newNode.Next = newNode
	}

	queue.size = queue.size + 1
}

func (queue *LinkQueue) Remove() string {
	queue.lock.Lock()
	queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty")
	}

	topNode := queue.root
	v := topNode.Value

	queue.root = topNode.Next

	queue.size = queue.size - 1

	return v
}

func (queue *LinkQueue) Size() int {
	return queue.size
}

func main() {
	queue := new(LinkQueue)
	queue.Add("a")
	queue.Add("b")
	queue.Add("c")
	fmt.Println(queue.Size())
	fmt.Println(queue.root.Value)
	//TODO: 链表队列遍历

	queue.Remove()
	fmt.Println(queue.Size())
}
