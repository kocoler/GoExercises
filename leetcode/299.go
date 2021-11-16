package main

import (
	"strconv"
)

func getHint(secret string, guess string) string {
	m := make(map[byte]int)

	for k := range secret {
		m[secret[k]] ++
	}

	a := 0
	b := 0
	visited := make(map[int]struct{})
	for k := range guess {
		if secret[k] == guess[k] {
			a ++
			m[secret[k]] --
			visited[k] = struct{}{}
		}
	}

	for k := range guess {
		if _, ok := visited[k]; !ok {
			if m[guess[k]] != 0 {
				b ++
				m[guess[k]] --
			}
		}
	}

	return strconv.Itoa(a) + "A" + strconv.Itoa(b) + "B"
}
