package main

import "fmt"

func main() {
	fmt.Println(permutation("abc"))
}

func permutation(s string) []string {
	record = make(map[string]bool)
	dfs([]byte(""), []byte(s))

	var res []string
	for k, _ := range record {
		res = append(res, k)
	}

	return res
}

var record map[string]bool
func dfs(now, next []byte) {
	fmt.Println(string(now), string(next))
	if len(next) == 0 {
		record[string(now)] = true
		return
	}
	for k, v := range next {
		newNow := make([]byte, len(now))
		copy(newNow, now)
		newNow = append(newNow, v)
		newNext := make([]byte, len(next))
		copy(newNext, next)
		fmt.Println(len(newNext))
		newNext = append(newNext[:k], newNext[k+1:]...)
		dfs(newNow, newNext)
	}
}
