package main

import "fmt"

func main() {
	fmt.Println(checkStraightLine([][]int{{2, 1}, {4, 2}, {6, 3}}))
}

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) < 3 {
		return true
	}

	var a, k, b float64

	a = float64(coordinates[1][0] - coordinates[0][0])
	k = float64(coordinates[1][1] - coordinates[0][1])
	// fmt.Println(a, k)
	if a == 0 {
		k = 1
	} else {
		k = k / a
	}
	// fmt.Println(a, k)
	b = float64(coordinates[1][1]) - k * float64(coordinates[1][0])
	// y = kx + b

	fmt.Println("y = ", k, "x", " + ", b)

	lenc := len(coordinates)
	for i := 2; i < lenc; i ++ {
		if ()
		x1 := float64(coordinates[i][0])
		y1 := float64(coordinates[i][1])
		x2 := float64(coordinates[i-1][0])
		y2 := float64(coordinates[i-1][1])
		x3 := float64(coordinates[i-2][0])
		y3 := float64(coordinates[i-2][1])
		a1 := (y1-y2)/(x1-x2)
		a2 := (y2-y3)/(x2-x3)
		if a1 == 0 {
			if a2 != 0 {
				return false
			}
		}
		if
		// fmt.Println("x = ", coordinates[i][0], "y = ", coordinates[i][1])
		// fmt.Println("y = ", k*coordinates[i][0]+b, "y = ", coordinates[i][1])
		if (k*float64(coordinates[i][0])+b) != float64(coordinates[i][1]) {
			return false
		}
	}

	return true
}
