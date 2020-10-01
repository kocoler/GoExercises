package main

import (
	"fmt"
	"strings"
)

/*使用函数包：
func FindSubstring(str1, str2 string) int {
	return strings.Index(str1,str2)
}*/

/* 未使用切片：
func FindSubstring(str1, str2 string) int {
	if(len(str1)<len(str2)) {
		return -1
	}
	n:=0
	flag:=1
	for i:=0;i<len(str1);i++  {
		if(str1[i]==str2[0]){
			n=1
			for j:=0;j<len(str2) ;j++  {
				if n==len(str2){
					n=i
					flag=0
					break
				}
				if str1[i+j]==str2[j] {
					n++
				}else {
						break
				}
			}
			if flag==0{
				break
			}
		}
	}
	if(flag==1){
		n=-1
	}
	return n
}*/
//使用切片：
func FindSubstring(str1, str2 string) int {
	if len(str1) < len(str2) {
		return -1
	}
	n := 0
	flag := 0
	for i := 0; i < len(str1); i++ {
		if len(str1)-i < len(str2) {
			break
		}
		if str1[i:i+len(str2)] == str2[0:len(str2)] {
			n = i
			flag = 1
			break
		}
	}
	if flag == 0 {
		n = -1
	}
	return n
}

func main() {
	var a1 string
	var a2 string
	fmt.Printf("输入两个字符串：")
	fmt.Scanf("%s", &a1)
	fmt.Scanf("%s", &a2)
	num := FindSubstring(a1, a2)
	fmt.Printf("%d", num)
}
