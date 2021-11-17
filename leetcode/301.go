package main

func isValid(str string) bool {
	cnt := 0
	for _, ch := range str {
		if ch == '(' {
			cnt++
		} else if ch == ')' {
			cnt--
			if cnt < 0 {
				return false
			}
		}
	}
	return cnt == 0
}

// map[string]struct{} 会更省空间
// BFS
func removeInvalidParentheses(s string) (ans []string) {
	curSet := map[string]bool{s: true}
	for {
		for str := range curSet {
			if isValid(str) {
				ans = append(ans, str)
			}
		}
		if len(ans) > 0 {
			return
		}
		nextSet := map[string]bool{}
		for str := range curSet {
			for i, ch := range str {
				if i > 0 && byte(ch) == str[i-1] {
					continue
				}
				if ch == '(' || ch == ')' {
					nextSet[str[:i]+str[i+1:]] = true
				}
			}
		}
		curSet = nextSet
	}
}
