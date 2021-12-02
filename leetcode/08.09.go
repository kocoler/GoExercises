package main

func generateParenthesis(n int) []string {
	var res []string
	// 左边现有的括号，右边现有的括号
	var dfs func(int, int)
	var now []byte
	dfs = func(l, r int) {
		if l == n && r == n {
			temp := make([]byte, len(now))
			copy(temp, now)
			res = append(res, string(temp))
			return
		}
		if l < n {
			now = append(now, '(')
			dfs(l+1, r)
			now = now[:len(now)-1]
		}
		if l > r {
			now = append(now, ')')
			dfs(l, r+1)
			now = now[:len(now)-1]
		}
	}
	dfs(0, 0)

	return res
}
