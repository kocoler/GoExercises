package main

import "sort"

type RangeFreqQuery struct {
	m map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	m := make(map[int][]int)
	for i, v := range arr {
		m[v] = append(m[v], i)
	}

	return RangeFreqQuery{m: m}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	var count int
	v := this.m[value]
	l := sort.Search(len(v), func(i int) bool {
		if v[i] >= left {
			return true
		}
		return false
	})
	r := sort.Search(len(v), func(i int) bool {
		if v[i] >= right {
			return true
		}
		return false
	})
	if r == len(v) || v[r] > right {
		r--
	}
	count = r - l + 1

	return count
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
