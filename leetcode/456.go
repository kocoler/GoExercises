package main

import "math"

type stackInt struct {
	val []int
	num int
}

func (s *stackInt) pop() int {
	res := s.val[s.num-1]
	s.val = s.val[:s.num-1]
	s.num--
	return res
}

func (s *stackInt) push(i int) {
	s.val = append(s.val, i)
	s.num++
}

func (s *stackInt) top() int {
	return s.val[s.num-1]
}

func (s *stackInt) length() int {
	return s.num
}

// have a peak
// max-stack
func find132pattern(nums []int) bool {
	lenn := len(nums)
	s := stackInt{}
	min := math.MinInt64
	for i := lenn - 1; i >= 0; i -- {
		if min > nums[i] {
			return true
		}
		if s.length() > 0 {
			if s.top() < nums[i] {
				for s.length() > 0 && s.top() < nums[i] {
					min = s.pop()
				}
				s.push(nums[i])
			} else {
				if min < nums[i] {
					s.push(nums[i])
				}
			}
		} else {
			s.push(nums[i])
		}
	}

	return false
}

func calMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}
