package main


// 更巧妙的做法:
// dp[i+1][dst] += dp[i][src] ... => dp[0][0] = 1 => ... + ...
// 经 k 次循环，dp[k][n-1]就是了

var dis int
var rm map[int][]int
var res int
func numWays(n int, relation [][]int, k int) int {
	rm = make(map[int][]int)

	res = 0
	dis = n - 1

	for _, v := range relation {
		rm[v[0]] = append(rm[v[0]], v[1])
	}

	dfs(0, k)

	return res
}

func dfs(now, left int) {
	if left == 0 {
		if now == dis {
			res ++
		}
		return
	}
	nextLeft := left - 1
	for _, v := range rm[now] {
		dfs(v, nextLeft)
	}
}

