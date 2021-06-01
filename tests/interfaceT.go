package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Stringer interface {
	String() string
}

func tt(i interface{}) {
	e := i.(func() Binary)
	fmt.Println(e().String())
	fmt.Println(reflect.TypeOf(e))
}

func tc() Binary {
	return Binary(10)
}

func t() {
	//var d func()
	tt(tc)
}

type Binary uint64

func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 2)
}

func (i Binary) Get() uint64 {
	return uint64(i)
}

func main() {
	b := Binary(200)
	s := Stringer(b)
	fmt.Println(s.String())

	t()
}