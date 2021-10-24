package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBeauty([][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}))
}

type sortInts [][]int

func (receiver sortInts) Less(i, j int) bool {
	if receiver[i][0] == receiver[j][0] {
		return receiver[i][1] < receiver[j][1]
	}

	return receiver[i][0] < receiver[j][0]
}

func (receiver sortInts) Swap(i, j int) {
	receiver[i], receiver[j] = receiver[j], receiver[i]
}

func (receiver sortInts) Len() int {
	return len(receiver)
}

func maximumBeauty(items [][]int, queries []int) []int {
	lenq := len(queries)
	res := make([]int, lenq)

	sort.Sort(sortInts(items))
	fmt.Println(items)
	max := -0x3f3f3f
	for k := range items {
		if items[k][1] > max {
			max = items[k][1]
		} else {
			items[k][1] = max
		}
	}
	fmt.Println(items, max)
	for k, num := range queries {
		index := sort.Search(len(items), func(j int) bool {
			return items[j][0] > num
		})
		index--
		fmt.Println("search", num, index)
		if index == len(items) {
			res[k] = max
		} else if items[index][0] > num {
			res[k] = 0
		} else {
			res[k] = items[index][1]
		}
	}

	return res
}
