package main

import "math/rand"

type Solution struct {
	ori []int
	shuffle []int
}

//  Fisher-Yates 洗牌算法

func Constructor(nums []int) Solution {
	shuffle := make([]int, len(nums))
	copy(shuffle, nums)
	return Solution{
		ori: nums,
		shuffle: shuffle,
	}
}


func (this *Solution) Reset() []int {
	for i := range this.ori {
		this.shuffle[i]= this.ori[i]
	}

	return this.ori
}


func (this *Solution) Shuffle() []int {

	for i := range this.shuffle {
		r := rand.Intn(len(this.shuffle) - i)
		r = r + i
		this.shuffle[i], this.shuffle[r] = this.shuffle[r], this.shuffle[i]
	}

	return this.shuffle
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
