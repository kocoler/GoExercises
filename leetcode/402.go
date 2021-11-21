package main

import "fmt"

func main() {
	fmt.Println(removeKdigits("10001", 4))
}

// 单调栈，需要能快速想到单调栈哦
func removeKdigits(num string, k int) string {
	lenn := len(num)
	if k >= lenn {
		return "0"
	}

	stack := []byte{num[0]}
	step := 0
	for i := 1; i < lenn; i++ {
		if step >= k || len(stack) == 0 || stack[len(stack)-1] < num[i] {
			stack = append(stack, num[i])
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > num[i] && step < k {
				stack = stack[:len(stack)-1]
				step++
			}
			stack = append(stack, num[i])
		}
	}

	stack = stack[:len(stack) + step - k]
	fmt.Println(string(stack), step, k)

	index := 0
	for i := 0; i < len(stack); i++ {
		if stack[i] == '0' {
			index++
		} else {
			break
		}
	}

	stack = stack[index:]

	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
