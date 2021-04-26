package main

import (
	"fmt"
	"strconv"
)

type stack struct {
	val []string
	num int
}

func (s *stack) pop() string {
	res := s.val[s.num-1]
	s.val = s.val[:s.num-1]
	s.num--
	return res
}

func (s *stack) push(i string) {
	s.val = append(s.val, i)
	s.num++
}

func (s *stack) popBottom() string {
	res := s.val[0]
	s.val = s.val[1:]
	s.num --
	return res
}

func (s *stack) pushBottom(i string) {
	s.val = append([]string{i}, s.val ...)
	s.num++
}

func (s *stack) top() string {
	return s.val[s.num-1]
}

func (s *stack) length() int {
	return s.num
}

func main() {
	fmt.Println(evalRPN([]string{"4","13","5","/","+"}))
}

func evalRPN(tokens []string) int {
	var num stack
	var operator stack
	// lene := len(tokens)
	for _, v := range tokens {
		// v := tokens[i]
		if isDigits(v) {
			num.push(v)
		} else {
			operator.push(v)
		}
	}

	fmt.Println(num.val, operator.val)

	for num.length() > 1 {
		num.push(cal(num.pop(), num.pop(), operator.popBottom()))
	}

	res, _ := strconv.Atoi(num.pop())

	return res
}

func isDigits(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return false
	}

	return true
}

func cal(a, b string, operation string) string {
	a, b = b , a
	aInt, _ := strconv.Atoi(a)
	bInt, _ := strconv.Atoi(b)
	var res int
	switch operation {
	case "*":
		res =  aInt * bInt
	case "-":
		res =  aInt - bInt
	case "+":
		res =  aInt + bInt
	case "/":
		res =  aInt / bInt
	}

	return strconv.Itoa(res)
}
