package main

import "fmt"

func main() {
	fmt.Println(numSubarrayProductLessThanK([]int{1,100,200,10,5,2,4,8,6,9}, 100))
}

// 乘积小于K的子数组
// 原来直接的 dp 超时了
// 很巧妙的双指针/滑动窗口: 子数组包含关系 和 一段数组中子数组的个数
// 还有一种二分解法
func numSubarrayProductLessThanK(nums []int, k int) int {
	lenn := len(nums)
	res := 0
	dpnum := 1
	// [1,100,200,10,5,2,4,8,6,9]
	// 1 2 4 5 6 8 9 10 100 200
	left := 0
	for i := 0; i < lenn; i ++ {
		dpnum *= nums[i]
		for dpnum >= k && left <= i {
			dpnum /= nums[left]
			left ++
		}
		res += i - left + 1
	}

	return res
}
