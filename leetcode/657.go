package main

func judgeCircle(moves string) bool {
	u := 0
	d := 0
	r := 0
	l := 0
	for _, v := range moves {
		switch v {
		case 'U':
			u++
		case 'D':
			d++
		case 'R':
			r++
		case 'L':
			l++
		}
	}

	return u == d && r == l
}
