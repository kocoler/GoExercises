package main

import "fmt"

type c struct {
	b int
	cs [10]*c
}

func construct() c {
	return c{
		b:  0,
		cs: [10]*c{},
	}
}

func main() {
	obj := construct()
	obj.cs[0] = &c{b: 0, cs: [10]*c{}}
	fmt.Println(obj.cs[0] == nil)
}
