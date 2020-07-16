package main

import "fmt"

func main() {
	// 创建一个容量为2的切片
	array := make([]int, 0, 2)
	array_info_print(array)

	_ = append(array, 1)
	array_info_print(array)
	_ = append(array, 1)
	array_info_print(array)
	_ = append(array, 1)
	array_info_print(array)

	fmt.Println("------")
	array = append(array, 1)
	array_info_print(array)
	array = append(array, 1)
	array_info_print(array)
	array = append(array, 1)
	array_info_print(array)
	array = append(array, 1, 1, 1, 1)
	array_info_print(array)
	array = append(array, 1, 1, 1, 1, 1, 1, 1, 1)
	array_info_print(array)
}

func array_info_print(array []int) {
	fmt.Println("cap", cap(array), "len", len(array), "array", array)
}
