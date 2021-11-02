package main

import "fmt"

func main() {
	fmt.Println(camelMatch([]string{"FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"}, "FB"))
}

// 双指针
// 没想的那么复杂 ...
// 可以更优的
// 在模式串后面加一个特殊字符，就少考虑了很多情况
// 大佬的写法是真的牛，思路清晰
func camelMatch(queries []string, pattern string) []bool {
	res := make([]bool, len(queries))
	lenp := len(pattern)

	for k, v := range queries {
		// p pointer
		p := 0
		// queries pointer
		q := 0
		res[k] = true
		lenq := len(v)
		for p < lenp && q < lenq {
			if pattern[p] == v[q] {
				p ++
				q ++
			} else {
				if isUpper(v[q]) {
					res[k] = false
					break
				}
				q ++
			}
		}
		if p != lenp {
			res[k] = false
		}

		if !res[k] || q >= lenq {
			continue
		}

		for i := q; i < lenq; i ++ {
			fmt.Println(i, q, lenq)
			if isUpper(v[i]) {
				res[k] = false
				break
			}
		}
	}

	return res
}

func isUpper(w byte) bool {
	if w < 'a' {
		return true
	}

	return false
}
