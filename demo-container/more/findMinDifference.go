package more

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

type StrHeap []string // 使用字符串切片来存储数据，并实现heap的相关接口

func (h StrHeap) Len() int           { return len(h) }
func (h StrHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h StrHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StrHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *StrHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findMinDifference(timePoints []string) int {
	h := StrHeap(timePoints)
	heap.Init(&h) // 初始化堆
	ret := 24 * 60
	pre := timeToInt(heap.Pop(&h).(string))
	first := pre
	next := 24 * 60
	for h.Len() > 0 {
		// 依次取出每个值，求差
		next = timeToInt(heap.Pop(&h).(string))
		tmp := next - pre
		if tmp > 12*60 {
			tmp = 24*60 - tmp
		}
		if tmp < ret {
			ret = tmp
		}
		pre = next
	}

	// 最值有可能他为首尾的差
	tmp := next - first
	if tmp > 12*60 {
		tmp = 24*60 - tmp
	}
	if tmp < ret {
		ret = tmp
	}

	return ret
}

// 将字符串时间转换为分钟数， 用以计算差值
func timeToInt(t string) int {
	fmt.Println(t)
	arr := strings.Split(t, ":")
	hour, _ := strconv.Atoi(arr[0])
	minute, _ := strconv.Atoi(arr[1])
	return hour*60 + minute
}
