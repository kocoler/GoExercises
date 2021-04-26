package main

import "fmt"

var dp [18][18]bool
var res [][]string
var n int
var str string

func partition(s string) [][]string {
	lens := len(s)
	n = lens
	s = fmt.Sprintf(" %s", s)
	str = s
	dp = [18][18]bool{}
	res = [][]string{}
	for i := lens; i >= 1; i-- {
		for j := i; j <= lens; j++ {
			if i == j {
				dp[i][j] = true
			}
			if s[i] == s[j] {
				if (i+1 == j) || dp[i+1][j-1] == true {
					dp[i][j] = true
				}
			}
		}
	}

	dfs(1)

	return res
}

var temp []string

func dfs(i int) {
	if i == n+1 {
		res = append(res, append([]string{}, temp...))
		return
	}

	for j := i; j <= n; j++ {
		if dp[i][j] {
			temp = append(temp, str[i:j+1])
			dfs(j + 1)
			temp = temp[:len(temp)-1]
		}
	}
}
