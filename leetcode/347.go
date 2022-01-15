package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func partition(nums [][2]int, l, r int) int {
	i := l - 1
	pivot := nums[r][1]
	for j := l; j < r; j++ {
		if nums[j][1] > pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[r], nums[i+1] = nums[i+1], nums[r]

	return i + 1
}

func randomPartition(nums [][2]int, l, r int) int {
	i := rand.Intn(r-l+1) + l

	nums[r], nums[i] = nums[i], nums[r]

	return partition(nums, l, r)
}

func qsort(nums [][2]int, l, r, k int) {
	if l < r {
		partition := randomPartition(nums, l, r)
		if partition == k {
			return
		}
		if partition > k {
			qsort(nums, l, partition-1, k)
		} else if partition < k {
			qsort(nums, partition+1, r, k)
		}
	}
}

func topKFrequent(nums []int, k int) []int {
	var occurs [][2]int
	occursMap := make(map[int]int)
	for _, v := range nums {
		occursMap[v]++
	}
	for k, v := range occursMap {
		occurs = append(occurs, [2]int{k, v})
	}
	qsort(occurs, 0, len(occurs)-1, k)
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = occurs[i][0]
	}
	return ret
}
