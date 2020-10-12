package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)
	var mute sync.Mutex
	go func() {
		for {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	for {
		mute.Lock()
		i += 1
		mute.Unlock()
	}
}
