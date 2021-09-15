package main

// 这个 db 比较神奇
// 他是对这个树的左右部分动归
// i 代表根结点
// j 代表左树长度

func numTrees(n int) int {
	num := make([]int, n+1)

	num[0] = 1
	num[1] = 1
	for i := 1; i <= n; i ++ {
		now := 0
		for j := 1; j <= i; j ++ {
			now = now + (num[j-1] * num[i-j])
		}
		num[i] = now
	}

	return num[n]
}
