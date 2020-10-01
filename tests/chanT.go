package main

import (
	"fmt"
	"sync"
)

func main() {
	num := []int{1,2,3,4,5}

	var wg sync.WaitGroup
	//var lock sync.Mutex
	errChan := make(chan error, 10)
	//defer close(errChan)
	for _, v := range num {
		wg.Add(1)
		v := v
		go func() {
			defer wg.Done()
			fmt.Println(v)
			if v==1 || v==5 {
				var err error
				errChan <- err
			}
			fmt.Println("wo hao le", v)
		}()
	}
	wg.Wait()
	close(errChan)
	panic("ddd")
	if len(errChan) != 0 {
		for v := range errChan {
			panic(v)
		}
	}
	fmt.Println("done")
	//fmt.Println("errChan:", len(errChan))
}
