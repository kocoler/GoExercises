package main

// 这个题比较有意义的就是 stack 中存的是索引！！！！！！！

func dailyTemperatures(temperatures []int) []int {
	lent := len(temperatures)
	res := make([]int, lent)
	var stack []int  // !!!!  这里存索引

	for i := 0; i < lent; i ++ {
		lens := len(stack)
		now := temperatures[i]
		for lens > 0 && temperatures[stack[lens - 1]] < now {
			res[stack[lens-1]] = i - stack[lens - 1]
			stack = stack[:lens-1]
			lens --
		}
		stack = append(stack, i)
	}

	return res
}
