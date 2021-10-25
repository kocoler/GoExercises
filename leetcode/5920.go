package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimizedMaximum(1, []int{1}))
}

// 周赛 T3，根本没想到二分 ...
// 这种二分我应该只见过一次？没有想到是这种 1e5 最大答案大范围的二分 ...
// 这里还有一种比较神奇的写法
// 比如 v % now ！= 0 ... 这一系列的做法可以直接简化成 v + now / now
// 判断 now == 0 的部分，也可以放到下面 (v + now) / (now + 1)
// 但是 now + 1 我看不懂的，为什么，最后答案还要 + 1，这两个操作是如何保证正确性的
func minimizedMaximum(n int, quantities []int) int {
	res := sort.Search(1e5, func(now int) bool {
		// check if now could ...
		num := 0
		if now == 0 {
			return false
		}

		for _, v := range quantities {
			num += v / now
			if v % now != 0 {
				num ++
			}
		}

		return num <= n
	})

	return res
}
