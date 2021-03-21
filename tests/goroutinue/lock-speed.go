package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2)
	withLock(1 << 22)
	withoutLock(1 << 22)
}

func productArray(n int) []int {
	var res []int
	for i := 0; i < n; i ++ {
		res = append(res, i)
	}

	return res
}

func withoutLock(n int) {
	ori := productArray(n)

	var res []int
	res = make([]int, len(ori))
	var sum int64

	var wg sync.WaitGroup

	t := time.Now()
	for k, v := range ori {
		k := k
		v := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			res[k] = v
		}()
	}

	wg.Wait()
	//for _, v := range res {
	//	sum += int64(v)
	//}

	fmt.Println("without lock:", time.Now().Sub(t))
	fmt.Println(sum)
}

func withLock(n int) {
	ori := productArray(n)

	var mute sync.Mutex
	var sum int64

	var wg sync.WaitGroup

	t := time.Now()
	for _, v := range ori {
		v := v
		wg.Add(1)
		go func() {
			defer wg.Done()
			mute.Lock()
			sum += int64(v)
			mute.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("with lock:", time.Now().Sub(t))
	fmt.Println(sum)
}
