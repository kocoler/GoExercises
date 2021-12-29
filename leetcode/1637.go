package main

import "sort"

// 天际线问题

// 不过好像是应该用桶排序来做？

type point [][]int

func (p point) Len() int {
	return len(p)
}

func (p point) Less(i, j int) bool {
	return p[i][0] < p[j][0]
}

func (p point) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func maxWidthOfVerticalArea(points [][]int) int {
	sort.Sort(point(points))

	res := 0

	for i := 1; i < len(points); i++ {
		res = max(res, points[i][0]-points[i-1][0])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
