package main

import "strings"

func wordBreak(s string, wordDict []string) []string {
	dp := make([]bool, len(s)+1)
	wm := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wm[v] = true
	}

	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wm[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	var res []string
	var backtrace func(index int)

	return res
}

// 先写备忘录
func wordBreak(s string, wordDict []string) []string {
	dp := make([][][]string, len(s))
	wm := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wm[v] = true
	}

	var res []string
	var backtrace func(index int) [][]string
	backtrace = func(index int) [][]string {
		if index >= len(s) {
			return [][]string{}
		}
		if dp[index] != nil {
			return dp[index]
		}

		var now [][]string
		for i := index + 1; i < len(s); i++ {
			word := s[index:i]
			if wm[word] {
				for _, v := range backtrace(i) {
					now = append(now, append([]string{word}, v...))
				}
			}
		}
		if wm[s[index:]] {
			now = append(now, []string{s[index:]})
		}

		dp[index] = now
		return now
	}

	for _, v := range backtrace(0) {
		res = append(res, strings.Join(v, " "))
	}

	return res
}
