package main

import "math"

// 记忆化搜索
func integerReplacement(n int) int {
	m := make(map[int]int)
	var dfs func(int) int

	dfs = func(now int) int {
		if now == 1 {
			return 0
		}
		if v, ok := m[now]; ok {
			return v
		}
		if now%2 == 0 {
			return dfs(now/2) + 1
		}
		return min(dfs(now-1), dfs(now+1)) + 1
	}

	return dfs(n)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// 位运算
func integerReplacement(n int) int {
	math.Pow
	res := 0
	for n != 1 {
		if n&1 == 0 {
			n >>= 1
		} else {
			if n == 3 {
				n -= 1
			} else if n&2 == 0 {
				// 2: 0xb10
				// 0bxxx01
				n--
			} else {
				// 0bxxx11
				n++
			}
		}
		res += 1
	}

	return res
}
