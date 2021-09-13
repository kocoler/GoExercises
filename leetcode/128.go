package main

// 哈希表
// 只更新头尾

func longestConsecutive(nums []int) int {
	m := make(map[int]int)
	res := 0

	for _, v := range nums {
		if m[v] != 0 {
			continue
		}
		// update head, update tail
		n := 1
		if m[v + 1] != 0 {
			n = n + m[v+1]
			m[v+n-1] = n
		}
		if m[v-1] != 0 {
			n = n + m[v-1]
			p := v - m[v-1]
			m[p] = n
			m[p+n-1] = n
		}
		m[v] = n
		if n > res {
			res = n
		}
	}

	return res
}
