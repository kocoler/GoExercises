package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(kMirror(3, 7))
}

// 1 2 3  + 1
// 11 22 33 44 55 + 11
// 111 121 131 141 151 + 010
// 212 222 232 + 100
// 1001 1111 1221  + 0110
// 66666 65556
func kMirror(k int, n int) int64 {
	var i, res int64

	i = 1
	for n > 0 {
		if verify(strconv.FormatInt(i, 10)) && verify(strconv.FormatInt(i, k)) {
			fmt.Println("verify!", i)
			res += i
			n--
		}
		i++
	}

	return res
}

func verify(str string) bool {
	l := 0
	r := len(str) - 1

	for l < len(str) {
		if str[l] != str[r] {
			return false
		}
		l++
		r--
	}

	return true
}
