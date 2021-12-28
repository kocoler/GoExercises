package main

import "fmt"

func main() {
	fmt.Println(maximalNetworkRank(5, [][]int{{2, 3}, {0, 3}, {0, 4}, {4, 1}}))
}

// 这个题有 O(m+n) 解法
// 就是再处理一次，把 first 和 second 找出来
// 再把 first 和 second list 找出来
// 如果 first 是 1，从 second 中找到不相连的，否则返回 first + second - 1
// 如果 first 不是 1，从 first 中找到不相连的，否则返回 first * 2 - 1

func maximalNetworkRank(n int, roads [][]int) int {
	// n <-> edges
	m := make(map[int]int)
	// edges
	medges := make([][]int, n)
	for k := range medges {
		medges[k] = make([]int, n)
	}
	for _, v := range roads {
		m[v[0]]++
		m[v[1]]++
		medges[v[0]][v[1]] = 1
		medges[v[1]][v[0]] = 1
	}

	var res int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			res = max(res, m[i]+m[j]-medges[i][j])
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
