package main

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for k := range nums {
		num := (nums[k] - 1) % n
		nums[num] += n
	}

	res := []int{}

	for k := range nums {
		if nums[k] <= n {
			res = append(res, k+1)
		}
	}

	return res
}
