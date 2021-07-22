package main

// 双指针

// 更优做法: 优化到时间复杂度为 1 => 从两边开始向中间，很神奇

func sortedSquares(nums []int) []int {
	lenn := len(nums)

	res := make([]int, lenn)
	l := 0
	r := 0
	for r < lenn && nums[r] < 0 {
		r ++
	}

	l = r - 1
	lenr := 0

	for r < lenn && nums[r] == 0 {
		res[lenr] = 0
		lenr ++
		r ++
	}

	var lv, lr int
	for ; l >= 0 && r < lenn; lenr ++ {
		lv = nums[l] * nums[l]
		lr = nums[r] * nums[r]
		if lv < lr {
			res[lenr] = lv
			l --
		} else {
			res[lenr] = lr
			r ++
		}
	}

	for r < lenn {
		res[lenr] = nums[r] * nums[r]
		r ++
		lenr ++
	}
	for  l >= 0 {
		res[lenr] = nums[l] * nums[l]
		l --
		lenr ++
	}

	return res
}
