package main

// 二分
// 本题是因为针对 mid 这个值，左边(<=)是要的，而右边(>)是不要的
// 因此需要 l = mid, r = mid - 1 右边缩小
// 因此要使 mid 靠右
// 也就是向上取整

func arrangeCoins(n int) int {
	l := 1
	r := n
	mid := 0
	for l < r {
		mid = l + (r - l + 1) / 2

		if (1 + mid) * mid / 2 <= n {
			l = mid
		} else {
			r = mid - 1
		}
	}

	return l
}
