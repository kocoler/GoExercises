package main

// 如果 minSize 有结果，那么 maxSize 一定有结果，所以可以不用考虑 maxSize
// 牛哇，我读题肯定读不出来

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	var res int
	m := make(map[string]int)
	lens := len(s)
	for i := 0; i <= lens-minSize; i++ {
		record := make(map[byte]struct{})
		for j := i; j < lens; j++ {
			now := s[i:j]
			record[s[j]] = struct{}{}
			if len(record) > maxLetters || len(now) > maxSize {
				break
			}
			if len(now) >= minSize {
				m[now]++
				if m[now] > res {
					res = m[now]
				}
			}
		}
	}

	return res
}

// 根据上面那个优化过的

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	var res int
	m := make(map[string]int)
	lens := len(s)
	for i := 0; i <= lens-minSize; i++ {
		record := make(map[byte]struct{})
		for j := i; j < i+minSize; j++ {
			record[s[j]] = struct{}{}
		}
		if len(record) > maxLetters {
			continue
		}
		now := s[i : i+minSize]
		m[now]++
		if m[now] > res {
			res = m[now]
		}
	}

	return res
}
