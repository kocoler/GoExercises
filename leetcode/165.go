package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1s := strings.Split(version1, ".")
	v2s := strings.Split(version2, ".")

	for i := 0; i < max(len(v1s), len(v2s)); i ++ {
		v1, v2 := 0, 0
		if i < len(v1s) {
			v1, _ = strconv.Atoi(v1s[i])
		}
		if i < len(v2s) {
			v2, _ = strconv.Atoi(v2s[i])
		}
		if v1 > v2 {
			return 1
		}
		if v1 < v2 {
			return -1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
