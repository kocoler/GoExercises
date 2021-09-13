package main

// 位运算
// 单独计算每个位 1 的数量
// 然后比较

func findDuplicate(nums []int) int {
	bitMax := 0
	lenn := len(nums)
	n := lenn - 1
	for n >> bitMax != 0 {
		bitMax ++
	}
	res := 0

	// 第 i 位
	for i := 0; i <= bitMax; i ++ {
		x := 0
		y := 0
		for j := 1; j <= n; j ++ {
			if j >> i & 1 == 1 {
				x ++
			}
		}

		for j := 0; j < lenn; j ++ {
			if nums[j] >> i & 1 == 1 {
				y++
			}
		}

		if x < y {
			res = res | (1 << i)
		}
	}

	return res
}
