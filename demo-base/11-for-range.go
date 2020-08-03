package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for index, value := range slice {
		// 因为for range创建了每个元素的副本，而不是直接返回每个元素的引用，
		// 如果使用该值变量的地址作为指向每个元素的指针，就会导致错误，
		// 在迭代时，返回的变量是一个迭代过程中根据切片依次赋值的新变量，
		// 所以值的地址总是相同的，导致结果不如预期。
		// https://studygolang.com/articles/9701
		// myMap[index] = &value
		num := value
		myMap[index] = &num
	}

	fmt.Println("=========new map============")
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}
