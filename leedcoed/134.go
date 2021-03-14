package main

func canCompleteCircuit(gas []int, cost []int) int {
	var max int
	var maxi int
	// var now int
	var uni int
	leng := len(gas)
	for i := leng - 1; i >= 0; i-- {
		r := gas[i] - cost[i]
		uni += r
		if uni >= max {
			maxi = i
			max = uni
		}
	}

	if uni < 0 {
		return -1
	}

	return maxi
}

// -2 -2 -2 3 3
// 1 -3 1 -2 3
