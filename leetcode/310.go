package main

import "fmt"

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
}

// 拓扑排序

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	// points := make(map[int][]int, n)
	points := make([][]int, n)
	// pointsNum := make(map[int]int, n)
	pointsNum := make([]int, n)
	for i := 0; i < len(edges); i++ {
		points[edges[i][0]] = append(points[edges[i][0]], edges[i][1])
		pointsNum[edges[i][0]]++
		points[edges[i][1]] = append(points[edges[i][1]], edges[i][0])
		pointsNum[edges[i][1]]++
	}

	var queue []int
	for k, v := range pointsNum {
		if v == 1 {
			queue = append(queue, k)
		}
	}

	last := n
	for len(queue) != last {
		lenq := len(queue)
		last -= lenq

		for i := 0; i < lenq; i++ {
			num := queue[i]
			for _, v := range points[num] {
				pointsNum[v]--
				if pointsNum[v] == 1 {
					queue = append(queue, v)
				}
			}
		}
		queue = queue[lenq:]
	}

	return queue
}
