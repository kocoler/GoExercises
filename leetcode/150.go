package main

import "strconv"

// 逆波兰表达式

func evalRPN(tokens []string) int {
	var numStack []int

	for i := 0; i < len(tokens); i++ {
		if isDigits(tokens[i]) {
			num, _ := strconv.Atoi(tokens[i])
			numStack = append(numStack, num)
		} else {
			num2, num1, res := numStack[len(numStack)-1], 0, 0
			if len(numStack) >= 2 {
				num1 = numStack[len(numStack)-2]
			}
			switch tokens[i] {
			case "*":
				res = num1 * num2
			case "/":
				res = num1 / num2
			case "-":
				res = num1 - num2
			case "+":
				res = num1 + num2
			}
			numStack = numStack[:len(numStack)-2]
			numStack = append(numStack, res)
		}
	}

	return numStack[0]
}

func isDigits(s string) bool {
	if s == "+" || s == "-" || s == "*" || s == "/" {
		return false
	}
	return true
}
