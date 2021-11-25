package main

import "math"

func main() {}

var res = []int{1, 2}
func constructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))

	var l int
	for w < area {
		if area % w == 0 {
			l = area / w
			break
		}
	}


	res[0] = w
	res[1] = l

	return res
}
