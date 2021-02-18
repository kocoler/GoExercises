package main

func mctFromLeafValues(arr []int) int {
	// dp
	// min-spanning-tree
	lena := len(arr)

	minStack := []int{0x3f3f3f}
	stackLen := 1
	res := 0
	for i := 0; i < lena; i ++ {
		// pop + push
		for arr[i] >= minStack[stackLen-1] {
			res += minStack[stackLen-1] * min(minStack[stackLen-2], arr[i])
			minStack = minStack[:stackLen-1]
			stackLen --
		}
		minStack = append(minStack, arr[i])
		stackLen ++
	}

	// stackLen > 2
	for stackLen > 2 {
		res += minStack[stackLen-1] * minStack[stackLen-2]
		minStack = minStack[:stackLen-1]
		stackLen --
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

