package main

import "fmt"

func main() {
	fmt.Println(getAverages([]int{8}, 1))
}

func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))
	if len(nums) < 2*k+1 {
		for k := range res {
			res[k] = -1
		}
		return res
	}

	sum := 0
	for i := 0; i < 2*k+1; i++ {
		sum += nums[i]
	}
	res[k] = sum / (2*k + 1)
	for i := k + 1; i < len(nums)-k; i++ {
		sum += nums[i+k]
		sum -= nums[i-k-1]
		res[i] = sum / (2*k + 1)
	}
	for i := 0; i < k; i++ {
		res[i] = -1
	}
	for i := len(nums) - k; i < len(nums); i++ {
		res[i] = -1
	}

	return res
}
