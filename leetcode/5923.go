package main

import "fmt"

// 简单的字符串处理

func main() {
	fmt.Println(minimumBuckets("H..H"))
}

func minimumBuckets(street string) int {
	var res int
	last := -2

	for i := 0; i < len(street); i++ {
		if street[i] == 'H' {
			if last != i-1 {
				if i+1 < len(street) && street[i+1] == '.' {
					last = i + 1
					res++
				} else if i-1 >= 0 && street[i-1] == '.' {
					last = i - 1
					res++
				} else if i+1 >= len(street) || street[i+1] == 'H' {
					return -1
				}
			}
		}
	}

	return res
}
