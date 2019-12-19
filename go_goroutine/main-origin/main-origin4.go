package main

import (
	"fmt"
	"runtime"
	"time"

	//"time"
)


func main() {
	var i int64 = -1
	runtime.GOMAXPROCS(2)
	//var mute sync.Mutex
	go func() {
		for {
			fmt.Println(i)
			time.Sleep(time.Second)
		}
	}()

	f := func(i *int64) {
		for {
			//mute.Lock()
			*i += 1
			//mute.Unlock()
		}
	}

	f(&i)

}