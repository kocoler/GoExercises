package main

import "time"

var A = 0
var B = 0

func a() {
	time.Sleep(time.Millisecond*5)
	B ++
	println("a", A)
}

func b() {
	time.Sleep(time.Millisecond*5)
	A ++
	println("b", B)
}

func main() {
	for i := 0; i < 10000; i ++ {
		go a()
		go b()
	}

	time.Sleep(time.Millisecond * 10)
}
