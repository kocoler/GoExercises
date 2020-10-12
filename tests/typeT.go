package main

import (
	"fmt"
	"strconv"
)

func main() {
	//a := time.Now()
	//fmt.Printf("%T",a)
	var a string = "111"
	b,_ := strconv.Atoi(string(a[1]))
	fmt.Println(b)
}
