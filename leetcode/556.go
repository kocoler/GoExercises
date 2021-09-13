package main

import (
	"sort"
	"strconv"
)

// 下一个元素
// 注意 max 32位的最大值
// 可以不用 sort.Slice 因为后面都是降序的

const max = int(^uint32(0)) >> 1

func nextGreaterElement(n int) int {
	s := strconv.AppendInt([]byte{}, int64(n), 10)
	lens := len(s)

	for i := lens - 1; i > 0; i-- {
		if s[i] > s[i-1] {
			n := s[i-1]
			sort.Slice(s[i-1:], func(j, k int) bool { return s[i-1+j] < s[i-1+k] })
			// fmt.Println()
			for j := i - 1; j < lens; j++ {
				if s[j] > n {
					for k := j; k > i-1; k-- {
						s[k], s[k-1] = s[k-1], s[k]
					}
					break
				}
			}
			break
		}
	}

	res, _ := strconv.Atoi(string(s))

	if res == n || res > max {
		return -1
	}

	return res
}
