package main

// 简单的滑动窗口
// 更优化一点不用每次都比较的方法是，维护一个 diff，记录每次滑动之后 不一样字符的个数
// 可以把每次两个数组比较的时间省掉

import "fmt"

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	mp := [26]int{}
	ms := [26]int{}
	res := []int{}

	for i := 0; i < len(p); i++ {
		mp[p[i]-'a']++
		ms[s[i]-'a']++
	}

	if ms == mp {
		res = []int{0}
	}
	for i := len(p); i < len(s); i++ {
		index := i - len(p) + 1
		ms[s[index-1]-'a']--
		ms[s[i]-'a']++
		if ms == mp {
			res = append(res, index)
		}
	}

	return res
}
