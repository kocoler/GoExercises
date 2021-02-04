package main

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

var st = struct {
	A int
	B int
}{
	A: 100,
	B: 100,
}

var i = with()

func with() int {
	return 2
}

func main() {
	sizeTest()
	// fmt.Println(i)
}

func sizeTest() {
	i := 1
	fmt.Println(binary.Size(struct{}{}))
	fmt.Println(unsafe.Sizeof(struct{}{}))
	fmt.Println(binary.Size([]byte{'1'}))
	fmt.Println(unsafe.Sizeof(i))
}
