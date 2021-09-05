package main

import "fmt"

func main() {
	a := []int{2}
	sortColors(a)
	fmt.Println(a)
}

func sortColors(nums []int) {
	p1 := 0
	lenn := len(nums)
	p2 := lenn - 1

	for i := 0; i <= p2; i++ {
		fmt.Println(i)
		if nums[i] == 0 {
			nums[p1], nums[i] = nums[i], nums[p1]
			p1++
		} else if nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
			i--
		}
	}
}
