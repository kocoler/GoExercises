package main

import (
	"fmt"
	"sync"
	"time"
)

func print(i int) {
	for i := 0; i <= 10; i ++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		print(1)
		wg.Done()
	}()
	for i := 10; i <= 20; i ++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
	wg.Wait()
}
