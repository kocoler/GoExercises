package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 拓扑排序
	in := make(map[int]int)    // 入度
	out := make(map[int][]int) // 出度
	queue := []int{}

	for i := 0; i < numCourses; i++ {
		in[i] = 0
	}

	// 0 <- 1
	for k := range prerequisites {
		incourse, outcourse := prerequisites[k][0], prerequisites[k][1]
		in[incourse]++
		out[outcourse] = append(out[outcourse], incourse)
	}

	for k, v := range in {
		if v == 0 {
			queue = append(queue, k)
		}
	}

	outnum := 0

	for len(queue) != 0 {
		lenq := len(queue)
		for i := 0; i < lenq; i++ {
			course := queue[i]
			// fmt.Println("seeing", course)
			outs := out[course]
			for _, v := range outs {
				in[v]--
				outnum := in[v]
				if outnum == 0 {
					queue = append(queue, v)
				}
			}
		}
		outnum += lenq

		queue = queue[lenq:]
	}

	// fmt.Println(outnum)
	// fmt.Println(in, out)

	if outnum == numCourses {
		return true
	}

	return false
}
