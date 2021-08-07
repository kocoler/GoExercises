package main

// 快慢双指针

var numsGlobal []int
var n int
func circularArrayLoop(nums []int) bool {
	n = len(nums)
	// lennGlobal = lenn
	numsGlobal = nums

	// m := make(map[int]bool)
	for i := 0; i < n; i ++ {
		if nums[i] == 0 {
			continue
		}
		slow := i
		fast := getNextStep(i)

		for nums[slow]*nums[fast] > 0 && nums[slow]*nums[getNextStep(fast)] > 0 {
			if slow == fast {
				if slow == getNextStep(slow) {
					break
				}
				return true
			}
			slow = getNextStep(slow)
			fast = getNextStep(getNextStep(fast))
		}

		// 这个是再次维护当前走过的路径
		// 也就是一段路走两边，打成维护两种状态的效果
		old := i
		for nums[old]*nums[getNextStep(old)] > 0 {
			tmp := old
			old = getNextStep(old)
			// 置 0， 利用原来的数组
			nums[tmp] = 0
		}
	}

	return false
}

func getNextStep(start int) int {
	start += numsGlobal[start]

	// 这个很神奇
	return ((start % n) + n) % n
}
