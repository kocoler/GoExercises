package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
}

func combinationSum3(k int, n int) [][]int {
	var res [][]int
	var dfs func(int, int)
	var now []int
	dfs = func(i, sum int) {
		sum += i

		if sum > n {
			return
		}
		// 这个位置注意一小下！
		if sum != n && len(now) >= k-1 {
			return
		}
		now = append(now, i)
		defer func() {
			now = now[:len(now)-1]
		}()
		if sum == n {
			if len(now) == k {
				temp := make([]int, len(now))
				copy(temp, now)
				res = append(res, temp)
			}
			return
		}
		for j := i + 1; j <= 9; j++ {
			dfs(j, sum)
		}
	}

	for i := 1; i <= 9; i++ {
		dfs(i, 0)
	}

	return res
}
