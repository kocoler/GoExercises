package main

import "fmt"

// de-queue
type stack struct {
	val []int
	num int
}

func (s *stack) pop() int {
	res := s.val[s.num-1]
	s.val = s.val[:s.num-1]
	s.num --
	return res
}

func (s *stack) popBottom() int {
	res := s.val[0]
	s.val = s.val[1:]
	s.num --
	return res
}

func (s *stack) push(i int) {
	s.val = append(s.val, i)
	s.num++
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
	fmt.Println(maxSlidingWindow([]int{1,-1}, 1))
}

func maxSlidingWindow(nums []int, k int) []int {
	lenn := len(nums)
	if lenn <= 0 {
		return nums
	}
	res := make([]int, lenn+2)

	var s stack
	// s.push(nums[0])
	left := 0
	right := -1
	for i := 0; i < lenn; i ++ {
		fmt.Println(s.val)
		if right - left >= k-1 {
			right ++
			fmt.Println(left, right)
			if s.bottom() == nums[left] {
				s.popBottom()
			}
			left ++
		} else {
			right ++
		}
		if s.length() > 0 {
			for s.length() > 0 && s.top() < nums[i] {
				s.pop()
			}
			s.push(nums[i])
			res[i] = s.bottom()
		} else {
			s.push(nums[i])
			res[i] = s.bottom()
		}
	}

	return res[k-1:lenn]
}
