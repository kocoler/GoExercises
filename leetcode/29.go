package main

import (
	"math"
)

// 需要考虑的情况比较多
// 都转负数可以表示的区间大？
// quick multi
// 注意负数的比较

// return z * y < x
func quickMulti(x, z, y int) bool {
	for sum := 0; z > 0; z >>= 1 {
		if z & 1 > 0 {
			sum += y

			// !! 这里都是负数！！
			if sum < x {
				return false
			}
		}

		if z != 1 {
			y += y
		}
	}
	return true
}

func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 { // 考虑被除数为最小值的情况
		if divisor == 1 {
			return math.MinInt32
		}
		if divisor == -1 {
			return math.MaxInt32
		}
	}
	if divisor == math.MinInt32 { // 考虑除数为最小值的情况
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	if dividend == 0 { // 考虑被除数为 0 的情况
		return 0
	}

	flag := false
	if dividend > 0 {
		dividend = -dividend
		flag = !flag
	}
	if divisor > 0 {
		divisor = -divisor
		flag = !flag
	}

	res := 0
	l, r := 0, math.MaxInt32
	for l <= r {
		mid := l + (r - l) >> 1
		if quickMulti(dividend, mid, divisor) {
			res = mid

			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if flag {
		return -res
	}
	return res
}
