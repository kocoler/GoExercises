package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(compress([]byte{'a', 'b', 'c', 'c', 'c', 'c', 'c', 'c'}))
}

func compress(chars []byte) int {
	n := 0
	p := chars[0]
	// chars = append(chars, ' ')
	lenc := len(chars)
	res := 0
	for i := 1; i < lenc; i++ {
		if p != chars[i] {
			n = n + 1
			if n == 1 {
				chars[res] = p
				res++
			} else if n < 10 {
				chars[res] = p
				r := itob(n)
				chars[res+1] = r[0]
				res += 2
			} else if n < 100 {
				chars[res] = p
				r := itob(n)
				chars[res+1] = r[0]
				chars[res+2] = r[1]
				res += 3
			} else if n < 1000 {
				chars[res] = p
				r := itob(n)
				chars[res+1] = r[0]
				chars[res+2] = r[1]
				chars[res+3] = r[2]
				res += 4
			}
			n = 0
			p = chars[i]
		} else {
			n++
		}
	}

	n = n + 1
	if n == 1 {
		chars[res] = p
		res++
	} else if n < 10 {
		chars[res] = p
		r := itob(n)
		chars[res+1] = r[0]
		res += 2
	} else if n < 100 {
		chars[res] = p
		r := itob(n)
		chars[res+1] = r[0]
		chars[res+2] = r[1]
		res += 3
	} else if n < 1000 {
		chars[res] = p
		r := itob(n)
		chars[res+1] = r[0]
		chars[res+2] = r[1]
		chars[res+3] = r[2]
		res += 4
	}

	// fmt.Println(string(chars))

	return res
}

func itob(i int) []byte {
	return strconv.AppendInt([]byte{}, int64(i), 10)
}
