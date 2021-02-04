package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Split("https://badu.com,https://github", ","))
	//fmt.Println(strings.Fields())
}
