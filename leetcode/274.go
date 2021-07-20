package main

// O(n)
// 其他解法: 基数排序

func hIndex(citations []int) int {
	sort.Ints(citations)

	lenc := len(citations)
	res := 0
	for i := 0 ; i < lenc; i ++ {
		if citations[i] >= (lenc - i) {
			res = min(citations[i], lenc - i)
			break
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
