package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	a := time.Now()
	b,_ := ioutil.ReadAll(a)
	fmt.Printf("%T",a)
}
