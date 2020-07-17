package main

import (
	"fmt"
	"sync"
)

type ArrayQueue struct {
	array []string
	size  int
	lock  sync.Mutex
}

func (queue *ArrayQueue) Add(v string) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	queue.array = append(queue.array, v)
	queue.size = queue.size + 1
}

func (queue *ArrayQueue) Remove() string {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	if queue.size == 0 {
		panic("empty")
	}

	v := queue.array[0]

	/*	for i := 1; i < queue.size; i++ {
			queue.array[i-1] = queue.array[i]
		}
		queue.array = queue.array[0 : queue.size-1]*/

	newArray := make([]string, queue.size-1, queue.size-1)
	for i := 1; i < queue.size; i++ {
		newArray[i-1] = queue.array[i]
	}
	queue.array = newArray

	queue.size = queue.size - 1
	return v
}

func (queue ArrayQueue) Size() int {
	return queue.size
}

func main() {
	queue := new(ArrayQueue)
	queue.Add("a")
	queue.Add("b")
	queue.Add("c")
	fmt.Println(queue.Size())
	fmt.Println(*queue)
	queue.Remove()
	fmt.Println(*queue)
	fmt.Println(queue.Size())
}
