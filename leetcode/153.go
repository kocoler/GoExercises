package main

var n []int
func findMin(nums []int) int {
	n = nums
	return binSearch(0, len(nums) - 1)
}

func binSearch(l, r int) int {
	if r - l <= 1 {
		if l == r {
			return n[l]
		} else {
			return min(n[l], n [r])
		}
	}

	mid := (l + r) / 2
	if n[l] < n[r] || (n[l] > n[mid] && n[mid] < n[r]){
		return binSearch(l, mid)
	} else {
		return binSearch(mid, r)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
