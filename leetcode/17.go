package main

import "fmt"

var res []string

func main() {
	fmt.Println(letterCombinations("23"))
}

var nummap []string = []string{
	"", // 0
	"", // 1
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {
	res = []string{}
	if len(digits) == 0 {
		return res
	}

	dfs("", digits)

	return res
}

func dfs(now string, aim string) {
	if len(aim) < 1 {
		res = append(res, now)
		return
	}

	n := aim[0] - 48
	lens := len(nummap[n])
	for i := 0; i < lens; i++ {
		// lenn := len(now)
		var next string
		//copy([]uint8(next), now)
		next = fmt.Sprintf("%s%s", now, string(nummap[n][i])) // append(next[j], string(nummap[n][i]))
		var nextAim string
		nextAim = fmt.Sprintf("%s", aim)
		// copy([]uint8(nextAim), []uint8(aim))
		// fmt.Println(nextAim)
		dfs(next, nextAim[1:])
	}

	return
}
