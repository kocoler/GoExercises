package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println(decodeCiphertext("iveo    eed   l te   olc", 4))
}

func decodeCiphertext(encodedText string, rows int) string {
	c := len(encodedText) / rows

	var bt bytes.Buffer
	// i -> c
	// j -> r
	for i := 0; i < c; i++ {
		num := min(rows, c-i)
		for j := 0; j < num; j++ {
			bt.WriteByte(encodedText[j*c+(j+i)])
		}
		//fmt.Println("now", bt.String(), num, i, c-i-1)
	}

	fmt.Println(c)

	return strings.TrimRight(bt.String(), " ")
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
