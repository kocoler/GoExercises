package main

import "fmt"

var dp [12]int
var status [12]bool // false -> - true -> +
var max int

// de-queue
type stack struct {
	val []int
	num int
}

func (s *stack) pop() int {
	res := s.val[s.num-1]
	s.val = s.val[:s.num-1]
	s.num--
	return res
}

func (s *stack) popBottom() int {
	res := s.val[0]
	s.val = s.val[1:]
	s.num--
	return res
}

func (s *stack) push(i int) {
	s.val = append(s.val, i)
	s.num++
}

func (s *stack) pushBottom(i int) {
	s.val = append([]int{i}, s.val...)
}

func (s *stack) top() int {
	return s.val[s.num-1]
}

func (s *stack) bottom() int {
	return s.val[0]
}

func (s *stack) length() int {
	return s.num
}

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}))
}

func wiggleMaxLength(nums []int) int {
	dp = [12]int{}
	status = [12]bool{}
	lenn := len(nums)
	max = 0
	for i := 1; i < lenn; i++ {
		now := nums[i] - nums[i-1]
		if i == 1 {
			dp[i-1] = 1
			dp[i] = 2
			if now > 0 {
				status[i] = true
			}
			// max = 2
			continue
		} else if now > 0 {
			status[i] = true
			if status[i-1] {
				// i-1 > 0
				dp[i] = 1
			} else {
				dp[i] = dp[i-1] + 1
				// max = calMax(max, dp[i])
			}
		} else {
			// now < 0
			if !status[i-1] {
				// i-1 < 0
				dp[i] = 1
			} else {
				dp[i] = dp[i-1] + 1
				// max = calMax(max, dp[i])
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(status)

	// lenn >= 2
	//flag := status[1] // +
	for i := 0; i < lenn; i++ {

	}

	return max
}

func calMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}
