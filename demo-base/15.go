package main

import "fmt"

// golang 判断key是否在map中的方法
// https://www.php.cn/be/go/438502.html

func main() {
	myMap := make(map[int64]string)

	myMap[1] = "value1"
	myMap[2] = "value2"
	myMap[5] = "value5"
	myMap[6] = "value6"

	for _, num := range []int64{1, 2, 3, 4, 5, 6} {
		if _, ok := myMap[num]; ok {
			fmt.Printf("myMap中包含key: %d\n", num)
		} else {
			fmt.Printf("myMap中不包含key: %d\n", num)
		}
	}

	fmt.Println("===================分割线=======================")

	for _, num := range []int64{1, 2, 3, 4, 5, 6} {
		if v, s := myMap[num]; s {
			fmt.Printf("myNap中包含key：%d, value值为： %s\n", num, v)
		} else {
			fmt.Printf("myMap中不包含key:%d\n", num)
		}
	}
}
