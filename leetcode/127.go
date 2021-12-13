package main

import "fmt"

func main() {
	fmt.Println(ladderLength("a", "c", []string{"a", "b", "c"}))
}

// 单向 BFS

//func ladderLength(beginWord string, endWord string, wordList []string) int {
//	// bfs
//	step := 1
//
//	queue := [][]byte{[]byte(beginWord)}
//	m := make(map[string]bool)
//	for _, v := range wordList {
//		m[v] = true
//	}
//	if !m[endWord] {
//		return 0
//	}
//	for len(queue) > 0 {
//		step++
//		lenq := len(queue)
//		for i := 0; i < lenq; i++ {
//			// 换
//			for k := 0; k < len(queue[i]); k++ {
//				for j := 0; j < 26; j++ {
//					ori := queue[i][k]
//					queue[i][k] = byte('a' + j)
//
//					if m[string(queue[i])] {
//						if string(queue[i]) == endWord {
//							fmt.Println(string(queue[i]), endWord, m[string(queue[i])])
//							return step
//						}
//						m[string(queue[i])] = false
//						// 入队
//						nb := make([]byte, len(queue[i]))
//						copy(nb, queue[i])
//						queue = append(queue, nb)
//					}
//					queue[i][k] = ori
//				}
//			}
//		}
//		queue = queue[lenq:]
//	}
//
//	return 0
//}

// 双向 BFS，但是没有快太多，可能是因为测试数据太弱了吧，但是已经可以看出来快了，那数据量更大的情况下，会更明显

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// bfs
	step := 0

	queue := [][]byte{[]byte(beginWord)}
	queueRev := [][]byte{[]byte(endWord)}
	m := make(map[string]bool)
	for _, v := range wordList {
		m[v] = true
	}
	if !m[endWord] {
		return 0
	}
	m[beginWord] = false
	m[endWord] = false

	cur := make(map[string]int)
	cur[beginWord] = 1
	other := make(map[string]int)
	other[endWord] = 1
	find := func(queue [][]byte, cur, other map[string]int, curstep int) [][]byte {
		lenq := len(queue)
		for i := 0; i < lenq; i++ {
			// 换
			for k := 0; k < len(queue[i]); k++ {
				for j := 0; j < 26; j++ {
					ori := queue[i][k]
					queue[i][k] = byte('a' + j)

					if m[string(queue[i])] {
						m[string(queue[i])] = false
						// 入队
						nb := make([]byte, len(queue[i]))
						copy(nb, queue[i])
						queue = append(queue, nb)
						cur[string(nb)] = curstep
					} else {
						if otherStep, ok := other[string(queue[i])]; ok {
							fmt.Println(string(queue[i]), endWord, m[string(queue[i])])
							step = curstep + otherStep
							return queue
						}
					}
					queue[i][k] = ori
				}
			}
		}
		return queue[lenq:]
	}
	curstep, otherstep := 1, 1
	for len(queue) > 0 || len(queueRev) > 0 {
		if len(queue) > len(queueRev) {
			queue = find(queue, cur, other, curstep)
			curstep++
		} else {
			queueRev = find(queueRev, other, cur, otherstep)
			otherstep++
		}
		if step != 0 {
			return step
		}
	}

	return step
}
