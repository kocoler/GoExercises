package main

func titleToNumber(columnTitle string) int {
	bit := 1
	res := 0

	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += bit * int(columnTitle[i]-'A'+1)
		bit *= 26
	}

	return res
}
