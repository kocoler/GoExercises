package main

// 双指针交换元素

func moveZeroes(nums []int)  {
	lenn := len(nums)

	l := 0
	for i := 0; i < lenn; i ++ {
		if nums[i] == 0 {
			l = i
			for i := l + 1; i < lenn; i ++ {
				if nums[i] != 0 {
					nums[l], nums[i] = nums[i], nums[l]
					l ++
				}
			}
			return
		}
	}
}
