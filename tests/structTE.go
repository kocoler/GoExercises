package main

import "fmt"

type t struct {
	f func(i int)
}

func main() {
	var w = t{
		f: func(i int) {
			fmt.Println(i)
		},
	}

	w.f(100)
}
