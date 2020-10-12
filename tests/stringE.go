package main

import (
	"fmt"
)

func main(){
	var s string
	s = "aaa"
	s = "a" + s
	fmt.Printf("%T",s)
}
