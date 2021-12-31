package main

import "bytes"

func convertToTitle(columnNumber int) string {
	// 28: AB
	bt := bytes.Buffer{}

	for columnNumber > 0 {
		columnNumber--
		bt.WriteByte(byte(columnNumber%26 + 'A'))
		columnNumber /= 26
	}

	return string(reverse(bt.Bytes()))
}

func reverse(b []byte) []byte {
	i := 0
	j := len(b) - 1
	for i < j {
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}

	return b
}
