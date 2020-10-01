package main

import "fmt"

func main() {
	var s string
	/*m := make(map[byte]int)
	fmt.Scanf("%s",&a)
	if a[0]=='w'{
		fmt.Printf("1")
	}
	m['w']=1
	m =make(map[byte]int)
	fmt.Printf("%d",m[a[0]])*/
	fmt.Scanf("%s", &s)
	m := make(map[byte]int)
	a := 1
	b := 0
	if len(s) == 0 {
		a = 0
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]] == 0 {
			b++
			m[s[i]] = i + 1
			//fmt.Printf("%d",b)
			if b > a {
				a = b
			}
		} else {
			if i+1 < len(s) {
				i = m[s[i]] - 1
				m = make(map[byte]int)
				b = 0
			}
		}

	}
	fmt.Printf("%d", a)
}

/*
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	a:=1;b:=0
	if(len(s)==0){
		a=0
	}
	for i := 0;i <len(s); i++ {
		if m[s[i]]!=1 {
			m[s[i]]=1
			b++
		}else{
			m := make(map[byte]int)
			m[s[i]]=1
			if b>a{
				a=b
			}
			b=1
		}
	}
	return a
}*/
