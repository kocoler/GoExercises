package main

import "fmt"

// 滑动窗口 或者 双指针
// 有更加简洁的写法，我的写法不够简洁
// 简洁的写法非常神奇，比如从负数开始填 ...

func main() {
	fmt.Println(checkInclusion("a", "ba"))
}

func checkInclusion(s1 string, s2 string) bool {
	s1m := [26]int{}
	s2m := [26]int{}
	s := 0
	f := 0

	for k := range s1 {
		s1m[s1[k]-'a']++
	}

	for f < len(s2) {
		if f-s == len(s1) {
			return true
		}
		for s <= f && s1m[s2[s]-'a'] == 0 {
			s++
		}

		b := s2[f] - 'a'
		if s2m[b] < s1m[b] {
			s2m[b]++
			f++
			continue
		}
		for s < f && s2[s] != s2[f] {
			if s2m[s2[s]-'a'] > 0 {
				s2m[s2[s]-'a']--
			}
			s++
		}
		if s < f && s2[s] == s2[f] && s2m[b] > 0 {
			s2m[b]--
			s++
		}
		f++
		if s2m[b] < s1m[b] {
			s2m[b]++
		}
	}

	if f-s == len(s1) {
		return true
	}

	return false
}
