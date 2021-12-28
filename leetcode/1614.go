package main

func maxDepth(s string) int {
	l := 0
	res := 0
	for _, v := range s {
		if v == '(' {
			l++
		} else if v == ')' {
			res = max(res, l)
			l--
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
