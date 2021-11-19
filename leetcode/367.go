package main

// 二分

func isPerfectSquare(num int) bool {
	l := 0
	r := num

	for l <= r {
		mid := l + (r - l) / 2
		n := mid * mid
		if n == num {
			return true
		} else if n < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
