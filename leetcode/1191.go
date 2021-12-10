package main

import "fmt"

func main() {
	fmt.Println(kConcatenationMaxSum([]int{-5, -2, 0, 0, 3, 9, -2, -5, 4}, 5))
}

func kConcatenationMaxSum(arr []int, k int) int {
	lena := len(arr)
	sres := 0
	cur := 0
	allsum := 0
	for i := 0; i < len(arr); i++ {
		cur = max(cur, 0) + arr[i]
		sres = max(sres, cur)
		allsum += arr[i]
	}

	if k == 1 {
		return sres
	}

	rightSum := arr[lena-1]
	rightMax := make([]int, lena)
	rightMax[lena-1] = arr[lena-1]
	for i := lena - 2; i >= 0; i-- {
		rightSum += arr[i]
		rightMax[i] = max(rightSum, rightMax[i+1])
	}

	spanSum := 0
	leftSum := 0
	leftMax := 0
	cur = 0
	for i := 0; i < lena; i++ {
		leftSum += arr[i]
		if i < lena-2 {
			spanSum = max(spanSum, leftSum+rightMax[i+2])
		}

		leftMax = max(leftMax, leftSum)
	}

	return max(allsum*k, max(sres, max(spanSum, leftMax+rightMax[0]+allsum*(k-2)))) % 1000000007
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
