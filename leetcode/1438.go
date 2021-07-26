package main

// 两个单调队列
func longestSubarray(nums []int, limit int) int {
	var minQ, maxQ []int
	res := 1

	minQ = append(minQ, nums[0])
	maxQ = append(maxQ, nums[0])

	lenn := len(nums)
	l := 0
	for i := 1; i < lenn; i ++ {
		value := nums[i]

		// lenMinQ := len(minQ)
		// lenMaxQ := len(maxQ)
		// for i := lenMinQ - 1; i >= 0; i -- {
		//     if minQ[i] <= value {
		//         break
		//     }
		//     minQ = minQ[:lenMinQ - 1]
		//     lenMinQ --
		// }
		// minQ = append(minQ, value)
		// for i := lenMaxQ - 1; i >= 0; i -- {
		//     if maxQ[i] >= value {
		//         break
		//     }
		//     maxQ = maxQ[:lenMaxQ - 1]
		//     lenMaxQ --
		// }
		// maxQ = append(maxQ, value)
		for len(minQ) > 0 && minQ[len(minQ)-1] > value {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, value)
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < value {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, value)
		for maxQ[0] - minQ[0] > limit {

			if nums[l] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[l] == maxQ[0] {
				maxQ = maxQ[1:]
			}

			l ++
		}
		res = max(res, i-l+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
