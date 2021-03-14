package main

import (
	"fmt"
	"strings"
)

//s = "1 + 1"
//输出：2
//1 <= s.length <= 3 * 105
//s 由数字、'+'、'-'、'('、')'、和 ' ' 组成
//s 表示一个有效的表达式

type stack struct {
	val []byte
	num int
}

func (s *stack) pop() byte {
	res := s.val[s.num-1]
	s.val = s.val[:s.num-1]
	s.num--
	return res
}

func (s *stack) push(i byte) {
	s.val = append(s.val, i)
	s.num++
}

func (s *stack) top() byte {
	return s.val[s.num-1]
}

func (s *stack) length() int {
	return s.num
}

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

func main() {
	fmt.Println(calculate("-2"))
}

func calculate(s string) int {
	lens := len(s)
	// s = s[5 : lens-1]
	s = strings.ReplaceAll(s, " ", "")
	s = s + "0"
	fmt.Println(s)

	var stackA stackInt // num
	var stackB stack // sign
	// var res int

	lens = len(s)
	for i := 0; i < lens; i++ {
		v := s[i]
		if isDigits(v) {
			// -
			num := 0
			for i < lens && isDigits(s[i]) {
				num = num*10 + int(s[i]-'0')
				i++
			}
			stackA.push(int(v - '0'))
			for stackB.length() > 0 && stackB.top() != '(' {
				stackA.push(cal(stackA.pop(), stackA.pop(), stackB.pop()))
				fmt.Println(stackA.val)
			}
			i --
		} else if v == '+' {
			stackB.push(v)
		} else if v == '-' {
			stackB.push(v)
		} else if v == '(' {
			stackB.push(v)
		} else if v == ')' {
			for stackB.length() > 0 && stackB.top() != '(' {
				stackA.push(cal(stackA.pop(), stackA.pop(), stackB.pop()))
			}
			if stackB.top() == '(' {
				stackB.pop()
			}
			for stackB.length() > 0 && stackB.top() != '(' {
				stackA.push(cal(stackA.pop(), stackA.pop(), stackB.pop()))
			}
		}
	}

	return stackA.pop()
}

func cal(a, b int, sign byte) int {
	if sign == '+' {
		return a + b
	}

	return b - a
}

func isDigits(b byte) bool {
	if 48 <= b && b <= 57 {
		return true
	}

	return false
}
