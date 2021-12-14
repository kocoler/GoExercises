package main

// 可以更细一点控制 index

func groupThePeople(groupSizes []int) [][]int {
	m := make(map[int][]int)
	var res [][]int

	for i := 0; i < len(groupSizes); i++ {
		m[groupSizes[i]] = append(m[groupSizes[i]], i)
		if len(m[groupSizes[i]]) == groupSizes[i] {
			res = append(res, m[groupSizes[i]])
			m[groupSizes[i]] = []int{}
		}
	}

	return res
}
