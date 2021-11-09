package main

import "strings"

func smallestSubsequence(s string, k int, letter byte, repetition int) string {
	stack := []byte{}
	if k == repetition {
		return strings.Repeat(string(letter), repetition)
	}

	num := 0
	for i := 0; i < lens; i ++ {
		if s[i] == letter {
			num ++
		}
	}
	outtime := num - repetition
	innum := 0
	for i := 0; i < lens; i ++ {
		// last letter must in right place
		if s[i] == letter {
			if {

			}
		}
		for len(stack) > 0 && stack[len(stack)-1] >= s[i] {
			if stack[len(stack) - 1] == letter {
				if outtime == 0 {
					break
				} else {

				}
			}
		}
		if len(stack) {

		}
	}


	return string(stack)
}
