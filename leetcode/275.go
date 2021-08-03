package main

// 二分，边界很重要

func hIndex(citations []int) int {
	lenc := len(citations)

	l := 0
	r := lenc - 1
	for l < r {
		mid := (l + r) / 2
		if citations[mid] >= (lenc - mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if citations[r] >= lenc - r {
		return lenc - r
	}
	return 0
}
