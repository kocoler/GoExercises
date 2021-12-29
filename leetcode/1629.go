package main

import (
	"math"
)

func slowestKey(releaseTimes []int, keysPressed string) byte {
	var res byte
	maxTime := math.MinInt64
	preTime := 0
	for i := 0; i < len(keysPressed); i++ {
		time := releaseTimes[i] - preTime
		if time >= maxTime {
			if time > maxTime || res < keysPressed[i] {
				res = keysPressed[i]
			}
			maxTime = time
		}
		preTime = releaseTimes[i]
	}

	return res
}
