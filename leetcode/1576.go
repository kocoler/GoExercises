package main

func modifyString(s string) string {
	sb := []byte(s)

	var pre byte
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			nx := byte('?')
			if i+1 < len(s) {
				nx = sb[i+1]
			}

			for k := 0; k < 26; k++ {
				b := byte('a' + k)
				if b != pre && b != nx {
					sb[i] = b
					break
				}
			}
		}
		pre = sb[i]
	}

	return string(sb)
}
