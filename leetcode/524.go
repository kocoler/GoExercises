package main

import (
	"sort"
)

// 可以优化的地方就是预先处理一下 dp[i][j] 表示 i 开始，第一个出现 j 的位置

func findLongestWord(s string, dictionary []string) string {
	sort.Slice(dictionary, func (n, m int) bool {
		return len(dictionary[n]) > len(dictionary[m]) || (len(dictionary[n]) == len(dictionary[m]) && dictionary[n] < dictionary[m])
	})

	lens := len(s)
	for _, v := range dictionary {
		f, t := 0, 0
		lenv := len(v)
		for f < lens {
			if s[f] == v[t] {
				f ++
				t ++
			} else {
				f ++
			}

			if t == lenv {
				return v
			}
		}
	}

	return ""
}
