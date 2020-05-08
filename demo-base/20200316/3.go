package main

import (
	"container/list"
	"fmt"
)

func main() {
	var initList list.List

	initList.PushBack(11)
	initList.PushBack(33)
	initList.PushBack(55)

	for element := initList.Front(); element !=  nil ; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

}
