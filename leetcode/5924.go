package main

// 贪心吧顶多算是

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	res := 0
	sx, sy := startPos[0], startPos[1]
	hx, hy := homePos[0], homePos[1]

	if sx < hx {
		for i := sx + 1; i <= hx; i++ {
			res += rowCosts[i]
		}
	}
	if sx > hx {
		for i := hx; i < sx; i++ {
			res += rowCosts[i]
		}
	}

	if sy < hy {
		for i := sy + 1; i <= hy; i++ {
			res += colCosts[i]
		}
	}

	if sy > hy {
		for i := hy; i < sy; i++ {
			res += colCosts[i]
		}
	}

	return res
}
