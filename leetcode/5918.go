package main

import "fmt"

func main() {
	fmt.Println(countVowelSubstrings("cuaieuouac"))
}

func countVowelSubstrings(word string) int {
	lenw := len(word)

	m := make(map[byte]int)
	num := 0
	res := 0
	lenn := 0
	for i := 0; i < lenw; i ++ {
		if word[i] == 'a' || word[i] == 'e' || word[i] == 'o' || word[i] == 'i' || word[i] == 'u' {
			if m[word[i]] == 0 {
				m[word[i]] = 1
				num ++
			} else {
				m[word[i]] ++
			}
			if num == 5 {
				res ++
				a, e, ii ,o , u := m['a'], m['e'], m['i'], m['o'], m['u']
				for j := i - lenn; j <= i; j ++ {
					fmt.Println(i, lenn, j)
					if m[word[j]] > 1 {
						res ++
						m[word[j]] --
					} else {
						break
					}
				}
				m['a'], m['e'], m['i'], m['o'], m['u'] = a, e, ii ,o , u
			}
			lenn ++
		} else {

			m['a'] = 0
			m['e'] = 0
			m['i'] = 0
			m['o'] = 0
			m['u'] = 0
			num = 0
			lenn = 0
		}
	}

	return res
}