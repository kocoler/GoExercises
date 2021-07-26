package main

import "sort"

// 因为 target 中元素不相同，所以可以以 target 中的元素做下标
// 用 arr 中元素做下标会出现其他的问题(多个下表)
func minOperations(target []int, arr []int) int {
	lent := len(target)
	lena := len(arr)
	d := []int{}
	m := make(map[int]int)
	for i := lent-1; i >= 0; i -- {
		m[target[i]] = i
	}

	lend := 0
	for i := 0; i < lena; i ++ {
		index := m[arr[i]]
		if target[index] == arr[i] {
			if lend == 0 || index > d[lend-1] {
				d = append(d, index)
				lend ++
			} else {
				idx := sort.SearchInts(d, index)
				d[idx] = index
			}
		}
	}

	return lent - lend
}
