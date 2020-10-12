package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)
	go func() {
		for {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	for {
		atomic.AddInt64(&i, 1)
	}

}
