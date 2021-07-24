package main

var dp [1001][1001]bool
func longestPalindrome(s string) string {
	dp = [1001][1001]bool{}
	lens := len(s)

	for i := 0; i < lens; i ++ {
		dp[i][i] = true
	}

	il := 0
	ir := 0
	m := 0

	for i := 2; i <= lens; i ++ {
		for j := 0; j+i-1 < lens; j ++ {
			r := j + i - 1
			flag := false
			if i > 3 {
				if dp[j+1][r-1] == true && s[j] == s[r] {
					dp[j][r] = true
					flag = true
				}
			} else {
				if s[j] == s[r] {
					dp[j][r] = true
					flag = true
				}
			}
			if flag && m < i {
				m = i
				il = j
				ir = r
			}
		}
	}


	return s[il:ir+1]
}

