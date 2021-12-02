package main

func twoSum(nums []int, target int) []int {
	var res []int
	m := make(map[int]int)

	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			res = []int{i, k}
			break
		}
		m[v] = k
	}

	return res
}
