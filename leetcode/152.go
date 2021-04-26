package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}

func maxProduct(nums []int) int {
	lenn := len(nums)
	if lenn < 1 {
		return 0
	}
	// res := make([]int, lenn+1)
	// res[0] = nums[0]
	max := nums[0]; min := nums[0]
	onMax := nums[0]; onMin := nums[0]

	for i := 1; i < lenn; i++ {
		//if res[i-1] < nums[i]*res[i-1] {
		//	res[i] = nums[i] * nums[i-1]
		//	continue
		//}
		//if res[i-1] < nums[i] {
		//	res[i] = nums[i]
		//	continue
		//}
		//res[i] = res[i-1]
		tempMax := maxInt(onMin*nums[i], onMax*nums[i], nums[i])
		fmt.Println(onMin*nums[i], onMax*nums[i], nums[i], "max:", tempMax)
		tempMin := minInt(onMin*nums[i], onMax*nums[i], nums[i])
		fmt.Println(onMin*nums[i], onMax*nums[i], nums[i], "min:", tempMin)
		if max < tempMax {
			max = tempMax
		}
		if min > tempMin {
			min = tempMin
		}
		onMax = tempMax
		onMin = tempMin
	}

	return max
}

func maxInt(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
	}
	if b > c {
		return b
	}

	return c
}

func minInt(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	}
	if b < c {
		return b
	}

	return c
}
