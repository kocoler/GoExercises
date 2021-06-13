package main

// 第一个错误的版本
// 二分

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var res int
// l = false
// r = true
var l int
var r int
func firstBadVersion(n int) int {
	l = 0
	r = n
	for r - l > 1 {
		mid := (l+r) / 2
		mr := isBadVersion(mid)
		if mr {
			r = mid
		} else {
			l = mid
		}
	}
	res = r

	return res
}

func bs(l, r int) {
	// fmt.Println(l, r)
	if r - l <= 1 {
		if !isBadVersion(l) && isBadVersion(r) {
			res = r
		}
		return
	}

	mid := (l+r) / 2
	mr := isBadVersion(mid)
	if mr {
		bs(l, mid)
	} else {
		bs(mid, r)
	}
}

