package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1,2,3}))
}

var lenn int
var res [][]int
var numsp []int
var m map[int]bool
func permute(nums []int) [][]int {
	lenn = len(nums)
	numsp = nums
	res = [][]int{}
	now = []int{}
	m = make(map[int]bool)

	for k, _ := range nums {
		backtrace(k)
	}

	return res
}

var now []int
func backtrace(position int) {
	m[numsp[position]] = true
	now = append(now, numsp[position])
	defer func() {
		m[now[len(now)-1]] = false
		now = now[:len(now)-1]
	}()
	fmt.Println(now, position)
	if len(now) == lenn {
		n := make([]int, len(now))
		copy(n, now)
		res = append(res, n)
		return
	}

	for i := 0; i < lenn; i ++ {
		// now = append(now, numsp[i+1])
		if !m[numsp[i]] {
			backtrace(i)
		}
		// now = now[:len(now)]
	}
}
