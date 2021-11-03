package main

import "fmt"

func main() {
	fmt.Println(smallestSubsequence("abcacababa"))
}

func smallestSubsequence(s string) string {
	lens := len(s)
	m := [26]int{}
	has := [26]int{}

	for i := 0; i < lens; i ++ {
		m[s[i] - 'a'] ++
	}

	// 最小栈
	stack := []byte{}
	for i := 0; i < lens; i ++ {
		m[s[i] - 'a'] --
		if has[s[i] - 'a'] == 1 {
			continue
		}
		//fmt.Println(string(s[i]), len(stack) - 1 > bottom, stack[len(stack) - 1] >= s[i])
		// 要注意维护的 m[stack[len(stack) - 1] - 'a'] > 0
		// 不要搞幺蛾子 呜呜呜
		for len(stack) > 0 && m[stack[len(stack) - 1] - 'a'] > 0 && stack[len(stack) - 1] >= s[i] {
			has[stack[len(stack)-1] - 'a'] = 0
			//fmt.Println("pop", string(stack[len(stack) - 1]))
			stack = stack[:len(stack) - 1]
		}
		has[s[i] - 'a'] = 1
		stack = append(stack, s[i])
	}

	return string(stack)
}
