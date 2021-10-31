package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findAllPeople(6, [][]int{{0, 2, 1}, {1, 3, 1}, {4, 5, 1}}, 1))
}

type meetingSort [][]int

func (m meetingSort) Len() int {
	return len(m)
}

func (m meetingSort) Less(i, j int) bool {
	return m[i][2] < m[j][2]
}

func (m meetingSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	m := make(map[int]struct{})
	sort.Sort(meetingSort(meetings))
	fmt.Println(meetings)

	m[0] = struct{}{}
	m[firstPerson] = struct{}{}
	res := []int{0, firstPerson}

	for i := 0; i < len(meetings); i++ {
		j := i
		for j < len(meetings)-1 && meetings[j+1][2] == meetings[i][2] {
			j++
		}

		relate := make(map[int][]int)
		for k := i; k <= j; k++ {
			relate[meetings[k][0]] = append(relate[meetings[k][0]], meetings[k][1])
			relate[meetings[k][1]] = append(relate[meetings[k][1]], meetings[k][0])
		}

		visited := make(map[int]bool)
		var update func(int)
		update = func(updatei int) {
			if visited[updatei] {
				return
			}
			visited[updatei] = true
			for _, v := range relate[updatei] {
				if _, ok := m[v]; !ok {
					m[v] = struct{}{}
					res = append(res, v)
					update(v)
				}
			}
		}

		for k := i; k <= j; k++ {
			//fmt.Println(k, meetings[k])
			if _, ok := m[meetings[k][0]]; ok {
				if _, ok := m[meetings[k][1]]; !ok {
					m[meetings[k][1]] = struct{}{}
					res = append(res, meetings[k][1])

					update(meetings[k][1])
				}
			}
			if _, ok := m[meetings[k][1]]; ok {
				if _, ok := m[meetings[k][0]]; !ok {
					m[meetings[k][0]] = struct{}{}
					res = append(res, meetings[k][0])

					update(meetings[k][0])
				}
			}
		}
		//fmt.Println(i, j, relate, m, res)

		i = j
	}

	return res
}
