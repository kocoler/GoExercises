package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isRectangleCover([][]int{{1, 2, 4, 4}, {1, 0, 4, 1}, {0, 2, 1, 3}, {0, 1, 3, 2}, {3, 1, 4, 2}, {0, 3, 1, 4}, {0, 0, 1, 1}}))
}

//[[1,2,4,4],[1,0,4,1],[0,2,1,3],[0,1,3,2],[3,1,4,2],[0,3,1,4],[0,0,1,1]]

type Line struct {
	x    int
	y1   int
	y2   int
	flag bool // true left, false right
}

type Lines []Line

func (l Lines) Len() int {
	return len(l)
}

func (l Lines) Less(i, j int) bool {
	if l[i].x == l[j].x {
		return l[i].y1 < l[j].y1
	}

	return l[i].x < l[j].x
}

func (l Lines) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

// 脑筋急转弯做法就是判断
// 总面积以及每个点的 2 / 4 个点
// 这里还是复习一下扫描线吧
func isRectangleCover(rectangles [][]int) bool {
	lenr := len(rectangles)
	lines := make(Lines, 2*lenr)
	idx := 0
	for k := range rectangles {
		x, y, a, b := rectangles[k][0], rectangles[k][1], rectangles[k][2], rectangles[k][3]

		lines[idx] = Line{x, b, y, true}
		idx++
		lines[idx] = Line{a, b, y, false}
		idx++
	}

	sort.Sort(lines)

	for l := 0; l < 2*lenr; {
		r := l
		var left, right, list Lines
		// 每次只扫描这条线的
		for r < 2*lenr && lines[r].x == lines[l].x {
			r++
		}

		for i := l; i < r; i++ {
			if lines[i].flag {
				list = left
			} else {
				list = right
			}

			if len(list) == 0 {
				list = append(list, lines[i])
			} else {
				if lines[i].y2 == list[len(list)-1].y1 {
					list[len(list)-1].y1 = lines[i].y1
				} else {
					list = append(list, lines[i])
				}
			}

			if lines[i].flag {
				left = list
			} else {
				right = list
			}
		}
		if l > 0 && r < 2*lenr {
			if len(left) != len(right) {
				return false
			}
			for k := range left {
				if left[k].y1 != right[k].y1 || left[k].y2 != right[k].y2 {
					return false
				}
			}
		} else {
			if len(left) != 1 && len(right) != 1 {
				return false
			}
		}

		l = r
	}

	return true
}
