package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println(regexp.Match("^[a-zA-Z0-9_]+$", []byte("1")))
}
