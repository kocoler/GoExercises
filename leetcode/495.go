package main

// 模拟

func findPoisonedDuration(timeSeries []int, duration int) int {
	lent := len(timeSeries)
	if lent == 0 {
		return 0
	}

	res := 0
	next := -1
	for i := 0; i < lent; i++ {
		res += duration
		if next >= timeSeries[i] {
			res -= (next - timeSeries[i] + 1)
		}

		next = timeSeries[i] + duration - 1
	}

	return res
}
