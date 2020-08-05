package main

import "fmt"

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		newIndex := abs(nums[i]) - 1

		if nums[newIndex] > 0 {
			nums[newIndex] *= -1
		}
	}

	var res []int

	for i := 1; i <= len(nums); i++ {
		if nums[i-1] > 0 {
			res = append(res, i)
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
