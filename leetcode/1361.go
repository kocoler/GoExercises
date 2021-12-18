package main

import "fmt"

func main() {
	fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}))
}

// 拓扑排序，自带过滤环的结构

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	in := make([]int, n)    // no more than 1 degree
	out := make([][]int, n) // no more than 2 degree

	initDegree := func(child []int) bool {
		for k, v := range child {
			if v == -1 {
				continue
			}
			in[v]++
			if in[v] > 1 {
				return false
			}
			out[k] = append(out[k], v)
			if len(out[k]) > 2 {
				return false
			}
		}
		return true
	}

	if !(initDegree(leftChild) && initDegree(rightChild)) {
		return false
	}

	var queue []int
	for k, v := range in {
		if v == 0 {
			queue = append(queue, k)
		}
	}
	if len(queue) != 1 {
		return false
	}

	onum := 1
	for len(queue) > 0 {
		lenq := len(queue)

		for i := 0; i < lenq; i++ {
			for _, v := range out[queue[i]] {
				in[v]--
				if in[v] == 0 {
					onum++
					queue = append(queue, v)
				}
			}
		}

		queue = queue[lenq:]
	}

	return onum == n
}
