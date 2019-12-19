package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var i int64 = -1
	var j int64
	ch := make(chan int64,1)
	runtime.GOMAXPROCS(2)
	ch <- i
	go func() {

		for {
			select {
			case j = <- ch:
				fmt.Println(j)
				//default:
			}
		time.Sleep(time.Second)
		}
	}()

	for {
		i += 1
		ch <- i
	}
}