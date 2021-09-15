package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
}

var str string
var lens int
func decodeString(s string) string {
	str = s
	lens = len(s)

	return recursion(0)
}

func recursion(index int) string {
	var bt bytes.Buffer

	num := 0

	for index < lens && str[index] != ']' {
		if isAlpha(str[index]) {
			bt.WriteByte(str[index])
		} else {
			for isNum(str[index]) {
				num = num * 10 + int(str[index] - '0')
				index ++
			}

			if num == 0 {
				num = 1
			}

			bt.WriteString(strings.Repeat(recursion(index + 1), num))

			l := 1
			r := 0
			index ++
			for index < lens {
				if str[index] == '[' {
					l ++
				}
				if str[index] == ']' {
					r ++
					if l == r {
						break
					}
				}
				index++
			}
			num = 0
		}
		index ++
	}

	return bt.String()
}

func isAlpha(b byte) bool {
	if b <= 'z' && b >= 'a' {
		return true
	}

	return false
}

func isNum(b byte) bool {
	if b <= '9' && b >= '0' {
		return true
	}

	return false
}
