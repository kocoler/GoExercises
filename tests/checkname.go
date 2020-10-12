package main

import (
	"log"
	"regexp"
)

func main() {
	str := "aA019Zzww22w"
	log.Println(CheckName(str))
}

func CheckName(str string) bool {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9_]{4,16}$", str); !ok {
		return false
	}
	return true
}

func CheckNamee(str string) bool {
	for _, v := range str {
		if (v < 123 && v > 96) || (v == 95) || (v > 64 && v < 91) || (v < 58 && v > 47) {
			continue
		} else {
			return false
		}
	}
	return true
}
