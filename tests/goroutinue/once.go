package main

import (
	"fmt"
	"sync"
	"time"
)

var so sync.Once

func do() {
	fmt.Println("befor do")
	time.Sleep(5 * time.Second)
	fmt.Println("after do")
}

func domain(i int) {
	fmt.Println(i, "doing")
	so.Do(do)
	fmt.Println(i, "done")
}

func main() {
	for i := 0; i < 10; i  ++ {
		domain(i)
	}
}
