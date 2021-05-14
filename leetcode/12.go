package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(intToRoman(58))
}

func intToRoman(num int) string {
	t := num
	b := 1
	var res []string
	for t > 0 {
		n := t % 10
		res = append([]string{calStr(n, b)}, res ...)
		b *= 10
		t = t / 10
	}

	return strings.Join(res, "")
}

// I II III IV V VI VII VIII IX X
func calStr(n int, b int) string {
	fmt.Println(n, b)
	p := n
	res := ""
	base := ""
	baseV := ""
	baseX := ""
	t := 0
	switch b {
	case 1:
		base = "I"
		baseV = "V"
		baseX = "X"
	case 10:
		base = "X"
		baseV = "L"
		baseX = "C"
	case 100:
		base = "C"
		baseV = "D"
		baseX = "M"
	case 1000:
		base = "M"
	}

	if p == 5 {
		return baseV
	}

	if p > 5 {
		if p == 9 {
			return base + baseX
		} else {
			t = p - 5
			res = baseV
		}
	} else {
		if p == 4 {
			return base + baseV
		} else {
			t = p
			res = ""
		}
	}

	for i := 0; i < t; i++ {
		res = res + base
	}

	return res
}
