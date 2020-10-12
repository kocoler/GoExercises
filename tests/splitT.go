package main

import (
	"fmt"
	"strings"
)

func main() {
	day, month := parseDumpTime("*/1:*")
	fmt.Println(day)
	fmt.Println(month)
}

func parseDumpTime(bumpTime string) (string, string) {
	tokens := strings.Split(bumpTime, ":")
	return tokens[0], tokens[1]
}
