package main

import (
	"container/list"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	var i int64 = -1
	l := list.New()
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(2)
	wg.Add(1)
	l.PushBack(i)
	go func() {

		for {
			if j := l.Front(); j != nil {
				fmt.Println(j.Value)
				l.Remove(j)
			}
			time.Sleep(time.Second)
		}
	}()

	for {
		i++
		l.PushBack(i)
	}

}
