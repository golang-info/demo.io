package main

import "fmt"

func main() {
	nums := []int{3, 4, 7, 8, 9}
	fmt.Println(maxProfit(nums))
}

func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	minprice := prices[0]
	maxprofit := 0
	for i := 1; i < length; i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		}

		// 第i天的最大收益 = max{前i - 1 天的最大收益， 第i天的价格 - 前 i -1 天中的最小价格}
		maxprofit = max(maxprofit, prices[i]-minprice)
	}
	return maxprofit
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
