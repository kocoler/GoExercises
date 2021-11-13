package main

// 摩尔投票还可以这么投哦
// 不止是半数以上

func majorityElement(nums []int) []int {
	num1, num2, vote1, vote2 := -0x3f3f3f, -0x3f3f3f, 0, 0
	lenn := len(nums)
	for i := 0; i < lenn; i ++ {
		if vote1 > 0 && num1 == nums[i] {
			vote1 ++
		} else if vote2 > 0 && num2 == nums[i] {
			vote2 ++
		} else if vote1 == 0 {
			num1 = nums[i]
			vote1 ++
		} else if vote2 == 0 {
			num2 = nums[i]
			vote2 ++
		} else {
			vote1 --
			vote2 --
		}
	}

	var res []int

	vote1, vote2 = 0, 0
	for i := 0; i < lenn; i ++ {
		if num1 == nums[i] {
			vote1 ++
		} else if num2 == nums[i] {
			vote2 ++
		}
	}

	if vote1 > lenn / 3 {
		res = append(res, num1)
	}
	if vote2 > lenn / 3 {
		res = append(res, num2)
	}

	return res
}
