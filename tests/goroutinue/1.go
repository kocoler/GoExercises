package main

import (
	"fmt"
	"sync"
)

func main() {
	a := 0
	b := 0

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		a = 1
		a++
	}()
	go func() {
		defer wg.Done()
		b = 10
		b--
	}()
	wg.Wait()

	fmt.Println(a)
	fmt.Println(b)
}
