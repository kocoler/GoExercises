package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var i int64 = -1
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(2)
	wg.Add(1)
	go func() {
		for {
			wg.Add(1)
			fmt.Println(i)
			wg.Done()
			time.Sleep(time.Second)

		}
	}()
	for {
		wg.Add(1)
		i++
		//time.Sleep(time.Second)
		wg.Done()
	}
}
