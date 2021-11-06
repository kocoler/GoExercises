package main

import "sort"

type bitsArray []int

func (b bitsArray) Len() int {
	return len(b)
}

func (b bitsArray) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b bitsArray) Less(i, j int) bool {
	bi := getBits(b[i])
	bj := getBits(b[j])
	if bi != bj {
		return bi < bj
	}

	return b[i] < b[j]
}

func sortByBits(arr []int) []int {
	sort.Sort(bitsArray(arr))
	return arr
}

func getBits(b int) int {
	res := 0
	for b > 0 {
		b -= n & (^n + 1)
	}

	return res
}
