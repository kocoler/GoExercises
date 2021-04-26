package main

import "fmt"

type NumArray struct {
	res []int
}

func Constructor(nums []int) NumArray {
	var res NumArray
	ln := len(nums)
	res.res = make([]int, ln+1)
	for i := 0; i < ln; i ++ {
		if i > 0 {
			res.res[i] = res.res[i-1] + nums[i]
		} else {
			res.res[i] = nums[i]
		}
	}
	res.res[0] = 0
	return res
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.res[right+1] - this.res[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */

func main() {
	// 0 -2 1 -4 -2 -3
	// -2 -2 1 -4 -2 -3 0
	Constructor([]int{-2, 0, 3, -5, 2, -1})
}
