package main

import "sort"

// 贪心
// 第 i 个位置，前 i - 1 个对当前是没有影响的
// 而根据当前的位置要预留出前面大数的留空位置
// 很巧妙

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || (people[i][0] == people[j][0] && people[i][1] > people[j][1])
	})

	lenp := len(people)
	res := make([][]int, lenp)

	for i := 0; i < lenp; i ++ {
		remain := 0
		index := 0
		for j := 0; j < lenp; j ++ {
			if res[j] == nil {
				remain ++
				if remain > people[i][1] {
					index = j
					break
				}
			}
		}
		res[index] = people[i]
	}

	return res
}
