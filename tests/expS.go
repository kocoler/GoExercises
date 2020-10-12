package main

import (
	"fmt"
	"regexp"
)

func main() {
	memberMap := make(map[string]string)
	str := `@kocoler`
	reg1 := regexp.MustCompile("@(.*?)[ \n\t]")
	reg2 := regexp.MustCompile("@([^ \n\t]*?)[^ \n\t]$")
	result2 := reg2.FindAllStringSubmatch(str, -1)
	for _, v := range result2 {
		v[1] += str[len(str)-1:]
		fmt.Println(v[1])
	}
	result1 := reg1.FindAllStringSubmatch(str,-1)
	for _, v := range result1 {
		fmt.Println(v[1])
	}
	//memberMap["ee"]="wee"
	//memberMap["eee"]="ww"
	for k, v := range memberMap {
		fmt.Println("kj")
		fmt.Println(k, v)
	}
}
