package main

import "fmt"

/*
	https://mp.weixin.qq.com/s/sTicB9IgCZxVhe53j7nsCA
*/

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 7
	fmt.Println(twoSum(nums, target)) // [1 4 2 3]
	fmt.Println(twoSumHash(nums, target))
}

func twoSum(nums []int, target int) []int {
	length := len(nums)
	ans := make([]int, 0)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				ans = append(ans, i, j)
			}
		}
	}
	return ans
}

/*
	官方采用hashmap的做法不就是你这种思路的优化版吗,保存一个hsahmap,
	然后遍历数组的同时看看当前元素的补数(与target的差值)是否存在于hashmap中,
	如果存在,就可以直接返回当前元素下标和hashmap中保存的下标,
	否则将当前元素的值和下标保存到hashmap中去

	list查找元素需要按照顺序一个一个查找,hashmap只需要查找一次就可以了
*/

func twoSumHash(nums []int, target int) []int {
	length := len(nums)
	result := make([]int, 0)
	m := make(map[int]int, length)
	for i, k := range nums {
		// 利用map中是否存在key为[target - k]的值
		if value, exist := m[target-k]; exist && value != i {
			// append 尾部追加元素
			result = append(result, value)
			result = append(result, i)
		}
		m[k] = i
	}
	return result
}
