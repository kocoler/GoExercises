package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}
	m := [26]int{}
	for i := 0; i < len(magazine); i++ {
		m[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		m[ransomNote[i]-'a']--
		if m[ransomNote[i]-'a'] < 0 {
			return false
		}
	}

	return true
}
