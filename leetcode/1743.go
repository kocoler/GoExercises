package main

var res []int
var m map[int][]int
func restoreArray(adjacentPairs [][]int) []int {
	m = make(map[int][]int)
	for _, pair := range adjacentPairs {
		p := pair[0]
		n := pair[1]
		m[p] = append(m[p], n)
		m[n] = append(m[n], p)
	}

	var init int
	for k, v := range m {
		if len(v) == 1 {
			init = k
			break
		}
	}

	lenm := len(m)
	res = make([]int, lenm)

	res[0] = init
	res[1] = m[init][0]
	for i := 2; i < lenm; i ++ {
		p := m[res[i-1]]
		if p[0] == res[i-2] {
			res[i] = p[1]
		} else {
			res[i] = p[0]
		}
	}

	return res
}
