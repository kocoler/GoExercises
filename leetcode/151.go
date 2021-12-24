package main

import "fmt"

// 好想调用 api 呀，哭

func main() {
	fmt.Println(reverseWords("good   example"))
}

func reverseWords(s string) string {
	bs := []byte(s)
	lens := len(s)

	trimRight := 0
	trimLeft := 0
	for i := lens - 1; i >= 0; i-- {
		if bs[i] == ' ' {
			trimRight++
		} else {
			break
		}
	}
	for i := 0; i < lens; i++ {
		if bs[i] == ' ' {
			trimLeft++
		} else {
			break
		}
	}
	bs = bs[trimLeft : lens-trimRight]
	fmt.Println(len([]byte("123 1")), len("123 1"))
	for i := 0; i < len(bs); i++ {
		if bs[i] == ' ' {
			fmt.Println("!")
			if bs[i-1] == ' ' {
				bs = append(bs[:i], bs[i+1:]...)
				i--
			}
			fmt.Println(string(bs))
		}
	}

	for i := 0; i < len(bs)/2; i++ {
		bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
	}

	for i := 0; i < len(bs); i++ {
		t := i
		for t = i; t < len(bs); t++ {
			if bs[t] == ' ' {
				break
			}
		}
		t--
		for k := 0; k < (t-i+1)/2; k++ {
			bs[i+k], bs[t-k] = bs[t-k], bs[i+k]
		}
		i = t + 1
	}

	return string(bs)
}
