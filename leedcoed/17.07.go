package main

import (
	"fmt"
	"regexp"
)

var reg = regexp.MustCompile("(.*?)\\((.*?)\\)")
var reg2 = regexp.MustCompile("\\((.*?),(.*?)\\)")

func main() {
	// reg := regexp.MustCompile(".*?\\((.*?)\\)")
	fmt.Println(reg.FindAllStringSubmatch("John(15)", -1)[0][1], reg.FindAllStringSubmatch("John(15)", -1)[0][2])
	fmt.Println(reg2.FindAllStringSubmatch("(Jon,John)", -1)[0][1], reg2.FindAllStringSubmatch("(Jon,John)", -1)[0][2])
	// strconv.Atoi()
}
