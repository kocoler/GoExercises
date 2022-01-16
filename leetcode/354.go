package main

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			return envelopes[i][1] < envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	dp := make([]int, len(envelopes))

	res := 1
	dp[0] = 1
	for i := 1; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][0] > envelopes[j][0] && envelopes[i][1] > envelopes[j][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

// 都忘了二分加速了 ...
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		// LIS 还是值得深入思考，这个地方的排序 一定是 >
		return envelopes[i][0] < envelopes[j][0] || envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})
	var stack [][]int

	for i := 0; i < len(envelopes); i++ {
		if index := sort.Search(len(stack), func(x int) bool {
			// >=
			return stack[x][0] >= envelopes[i][0] || stack[x][1] >= envelopes[i][1]
		}); index < len(stack) {
			stack[index] = envelopes[i]
		} else {
			stack = append(stack, envelopes[i])
		}
	}

	return len(stack)
}

// 优，都可以优
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		// LIS 还是值得深入思考，这个地方的排序 一定是 >
		return envelopes[i][0] < envelopes[j][0] || envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1]
	})

	var stack []int
	for i := 0; i < len(envelopes); i++ {
		if index := sort.Search(len(stack), func(x int) bool {
			// >=
			return stack[x] >= envelopes[i][1]
		}); index < len(stack) {
			stack[index] = envelopes[i][1]
		} else {
			stack = append(stack, envelopes[i][1])
		}
	}

	return len(stack)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
