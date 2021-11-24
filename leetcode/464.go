package main

import "fmt"

func main() {
	fmt.Println(canIWin(10, 11))
}

// 很多 贪心 都不对，需要完整的遍历所有的情况才行
// 所以这是个 dfs 的题
// 需要再加上 备忘录 提高性能
// 最开始想的 01 背包 + 贪心，后来发现是行不通的 ...
// 那也简单描述一下吧
// dp[i] 表示 能不能选择到 i 这个状态
// 如果某一轮选择到了 dp[desiredTotal]，那么判断当前 step 是否是 偶数，如果是就返回 true
// 否则返回 false
// 日后还是值得在写一遍的
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	// dp[i] 代表 i 的时候 输还是赢
	// -1 是输，1是赢，0是还未刷新状态
	dp := make([]int, 1<<maxChoosableInteger)

	var dfs func(record, added int) bool

	// record 记录了 哪一个 num 还没有被使用
	// 因为这道题说了 maxChoosableInteger 不会超过 30
	// ！这个优化也要记住！
	dfs = func(record, added int) bool {
		if dp[record] != 0 {
			if dp[record] == -1 {
				return false
			}
			return true
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			if record>>(i-1)&1 == 0 {
				// !dfs(record...) 就是说下一个步骤，也就是对方的输赢
				if i+added >= desiredTotal || !dfs(record|1<<(i-1), added+i) {
					dp[record] = 1
					return true
				}
			}
		}
		dp[record] = -1

		return false
	}

	return dfs(0, 0)
}
