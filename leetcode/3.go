package main

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	lens := len(s)
	res := 0
	l := 0
	r := 0
	for r < lens {
		v := s[r]
		if m[v] != 0 {
			for s[l] != v {
				m[s[l]] --
				l ++
			}
			l ++
		} else {
			m[v] ++
		}
		r ++

		if res < (r - l) {
			res = r - l
		}
	}

	return res
}
