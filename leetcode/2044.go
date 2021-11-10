package main

import "fmt"

func main() {
	fmt.Println(countMaxOrSubsets([]int{3, 2, 1, 5}))
}

var res int
var max int
var numsG []int

func countMaxOrSubsets(nums []int) int {
	res = 0
	max = 0
	numsG = nums
	lenn = len(nums)
	for k := range nums {
		max |= nums[k]
	}

	dfs(0, 0)

	return res
}

var lenn int

func dfs(index int, num int) {
	if index > lenn-1 {
		fmt.Println(index, num, max)
		if num == max {
			res++
		}
		return
	}

	nextnum := num | numsG[index]
	fmt.Println(num, numsG[index], nextnum)
	dfs(index+1, nextnum)
	dfs(index+1, num)
}
