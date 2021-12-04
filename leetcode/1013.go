package main

func canThreePartsEqualSum(arr []int) bool {
	lena := len(arr)
	preSum := make([]int, lena+1)

	for i := 1; i <= lena; i++ {
		preSum[i] = preSum[i-1] + arr[i-1]
	}
	sum := preSum[lena]

	if sum%3 != 0 {
		return false
	}

	for l, r := 0, lena-1; l < r-1; {
		lnum := preSum[l+1]
		mnum := preSum[r] - preSum[l+1]
		rnum := preSum[lena] - preSum[r]
		if lnum == mnum && mnum == rnum {
			return true
		}
		// 因为有负数，所以只要不等于，那就加
		if lnum != sum/3 {
			l++
		}
		if rnum != sum/3 {
			r--
		}
	}

	return false
}
