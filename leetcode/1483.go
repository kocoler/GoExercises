package main

// 倍增，LCA
// 2^j 倍的增

type TreeAncestor struct {
	dp [][]int // dp[node][j] -> dp[node] 的 2^j 个祖先，倍增
	// dp[node][j] = dp[dp[node][j-1]][j-1]
}

func Constructor(n int, parent []int) TreeAncestor {
	dp := make([][]int, n)
	for k, v := range parent {
		dp[k] = append(dp[k], v) // dp[k][0]
	}
	for j := 1; ; j++ {
		allneg := true

		for i := 0; i < n; i++ {
			// 这样可以不用考虑那么多边界，所有的 dp 一样长也是挺好的
			t := -1
			if dp[i][j-1] != -1 {
				t = dp[dp[i][j-1]][j-1]
			}
			dp[i] = append(dp[i], t)
			if t != -1 {
				allneg = false
			}
		}

		if allneg {
			break
		}
	}
	return TreeAncestor{dp: dp}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	res, pos := node, 0
	for k > 0 && res != -1 {
		if pos >= len(this.dp[res]) {
			res = -1
			break
		}
		if k&1 == 1 {
			res = this.dp[res][pos]
		}
		k >>= 1
		pos++
	}

	return res
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
