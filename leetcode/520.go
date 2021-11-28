package main

func detectCapitalUse(word string) bool {
	preUpper := 0
	for i := 0; i < len(word); i++ {
		if isUpper(word[i]) {
			if i > 0 && preUpper != i {
				return false
			}
			preUpper++
		} else if preUpper > 1 {
			return false
		}
	}

	return true
}

func isUpper(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}

	return false
}
