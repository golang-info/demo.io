package main

/*
	https://yougg.github.io/2017/06/12/go%E8%AF%AD%E8%A8%80%E5%AE%89%E5%85%A8%E7%BC%96%E7%A8%8B%E8%A7%84%E8%8C%83/#
	建议 2.4 慎用slice作为函数入参
*/

import (
	"fmt"
)

// slice 作为函数入参时时地址传递
func modify(array []int) {
	array[0] = 10  // [错误] 对入参 slice 的元素修改会影响到调用者
}

func main() {
	array := []int{1, 2, 3, 4, 5}

	modify(array)
	fmt.Println(array) // [输出]: [10, 2, 3, 4, 5]
}


/*

// 数组 作为函数入参时时地址传递
func modify(array [5]int) {
	array[0] = 10
}

func main() {
    // 【修改】 出入数组， 注意数组与slice的区别
	array := [5]int{1, 2, 3, 4, 5}

	modify(array)
	fmt.Println(array) // [输出]: [10, 2, 3, 4, 5]
}


*/