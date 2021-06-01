package main

import (
	"fmt"
	"plugin"
)

type Driver interface {
	Name() string
}

func main() {
	p, err := plugin.Open("driver.so")
	if err != nil {
		panic(err)
	}

	newDriverSymbol, err := p.Lookup("NewDriver")
	if err != nil {
		panic(err)
	}

	newDriverFunc := newDriverSymbol.(func() Driver)
	newDriver := newDriverFunc()
	fmt.Println(newDriver.Name())
}
