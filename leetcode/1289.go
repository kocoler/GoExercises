package main

func minFallingPathSum(arr [][]int) int {
	lenr := len(arr)
	lenc := len(arr[0])

	// dp := make([]int, lenc)

	// for i := 0; i < lenc; i ++ {
	//     dp[i] = arr[0][i]
	// }

	for i := 1; i < lenr; i ++ {
		for j := 0; j < lenc; j ++ {
			num := 0x3f3f3f
			for k := 0; k < lenc; k ++ {
				if k != j {
					num = min(num, arr[i-1][k])
				}
			}
			arr[i][j] = num + arr[i][j]
		}
	}

	res := 0x3f3f3f

	for _, v := range arr[lenr-1] {
		res = min(res, v)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

