package main

import "sort"

// 这思路真 nb

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] < properties[j][1]
		}
		return properties[i][0] > properties[j][0]
	})

	var res int
	var maxDef int
	for _, v := range properties {
		if v[1] < maxDef {
			res++
		} else {
			maxDef = v[1]
		}
	}

	return res
}
