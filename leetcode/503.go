package main

import "fmt"

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

func (s *stack) push(i int) {
	s.val = append(s.val, i)
	s.num++
}

func (s *stack) top() int {
	return s.val[s.num-1]
}

func (s *stack) length() int {
	return s.num
}

func main() {
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
}
// 1 2 1 1 2 1
func nextGreaterElements(nums []int) []int {
	var s stack
	temp := append(nums, nums ...)
	n := len(temp)
	res := make([]int, n)
	for i := n-1; i >= 0; i -- {
		fmt.Println(s.val, res)
		if s.length() > 0 {
			if s.top() > temp[i] {
				res[i] = s.top()
				s.push(temp[i])
			} else {
				var flag bool
				for s.length() > 0 {
					if s.top() > temp[i] {
						res[i] = s.top()
						s.push(temp[i])
						flag = true
						break
					} else {
						s.pop()
					}
				}
				if !flag {
					s.push(temp[i])
					res[i] = -1
				}
			}
		} else {
			s.push(temp[i])
			res[i] = -1
		}
	}

	return res[:n/2]
}
