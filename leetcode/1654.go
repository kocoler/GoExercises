package main

// BFS

func minimumJumps(forbidden []int, a int, b int, x int) int {
	if x == 0 {
		return 0
	}

	visited := make(map[[2]int]bool)
	visited[[2]int{0, 0}] = true
	forbiddenM := make(map[int]bool, len(forbidden))
	for _, v := range forbidden {
		forbiddenM[v] = true
	}

	queue := [][2]int{{0, 0}}

	step := 0
	for len(queue) > 0 {
		step++
		lenq := len(queue)
		for i := 0; i < lenq; i++ {
			now := queue[i]
			forward, backward := [2]int{now[0] + a, 0}, [2]int{now[0] - b, now[1] + 1}
			if backward[1] == 2 {
				backward[0] = 0
			}
			if backward[0] == x || forward[0] == x {
				return step
			}
			if !forbiddenM[forward[0]] && !visited[forward] && (forward[0] <= 8000) {
				visited[forward] = true
				queue = append(queue, forward)
			}
			if !forbiddenM[backward[0]] && backward[0] > 0 && !visited[backward] {
				visited[backward] = true
				queue = append(queue, backward)
			}
		}
		queue = queue[lenq:]
	}

	return -1
}
