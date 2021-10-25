package main

func countVowels(word string) int64 {
	var res int64
	lenw := len(word)

	for k := range word {
		if word[k] == 'a' || word[k] == 'e' || word[k] == 'i' || word[k] == 'o' || word[k] == 'u' {
			res += int64(lenw - k) * int64(k + 1)
		}
	}

	return res
}
