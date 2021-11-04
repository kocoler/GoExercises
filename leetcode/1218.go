package main

func longestSubsequence(arr []int, difference int) int {
	res := 1

	lena := len(arr)
	m := make(map[int]int)

	// dp[0] = 1
	for i := 0; i < lena; i ++ {
		m[arr[i]] = m[arr[i] - difference] + 1
		res = max(res, m[arr[i]])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
