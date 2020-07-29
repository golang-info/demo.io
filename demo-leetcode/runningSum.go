package main

/*
	https://mp.weixin.qq.com/s/fyzdzqDJPwuivo1OjKRvaw
*/

import "fmt"

func main() {
	// var demo []int
	demo := []int{1, 2, 3, 4, 5}
	fmt.Println(runningSum(demo))
}

func runningSum(nums []int) []int {
	length := len(nums)

	res := make([]int, length)

	if length < 1 {
		return res
	}

	res[0] = nums[0]
	for i := 1; i < length; i++ {
		res[i] = res[i-1] + nums[i]
	}
	return res
}
