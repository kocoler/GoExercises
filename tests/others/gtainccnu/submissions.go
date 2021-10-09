package main

import "fmt"

func main() {
	var length int
	_, err := fmt.Scanf("%d", &length)
	if err != nil {
		panic(err)
	}
	input := make([]int, length)
	for k, _ := range input {
		nownum := 0
		_, err = fmt.Scanf("%d", &nownum)
		if err != nil {
			panic(err)
		}
		input[k] = nownum
	}
	permute(input)

	fmt.Println(res)
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

	if len(now) == lenn {
		n := make([]int, len(now))
		copy(n, now)
		res = append(res, n)
		return
	}

	for i := 0; i < lenn; i ++ {
		if !m[numsp[i]] {
			backtrace(i)
		}
	}
}

