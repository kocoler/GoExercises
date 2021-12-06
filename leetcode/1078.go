package main

import "strings"

func findOcurrences(text string, first string, second string) []string {
	var res []string

	words := strings.Split(text, " ")

	same := first == second

	st := 0
	for i := 0; i < len(words); i++ {
		switch st {
		case 0:
			if words[i] == first {
				st = 1
			}
		case 1:
			if words[i] == second {
				st = 2
			} else if words[i] == first {
				st = 1
			} else {
				st = 0
			}
		case 2:
			res = append(res, words[i])
			if words[i] == second && same {
				st = 2
			} else if words[i] == first {
				st = 1
			} else {
				st = 0
			}
		}
	}

	return res
}
