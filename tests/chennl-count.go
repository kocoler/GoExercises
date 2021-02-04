package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	generator := func(done <-chan interface{}, start, max int, limit int) <-chan int {
		offsetStream := make(chan int)
		go func() {
			defer close(offsetStream)
			offset := start
			count := 0
			for count < max {
				// time.Sleep(time.Second)
				select {
				case <-done:
					return
				case offsetStream <- offset:
				}
				count += limit
				offset += limit
			}
		}()
		return offsetStream
	}

	getInfo := func(i int) error {

		time.Sleep(1*time.Second)
		if i == 400 {
			return errors.New("")
		}
		return nil
	}

	//resStream := func(done <- chan interface{}, errChan chan error, offsetStream <-chan int, limit int) <-chan int {
	//	resStream := make(chan int, 10)
	//	go func() {
	//		defer close(resStream)
	//		for i := range offsetStream {
	//			go func(i int) {
	//				err := getInfo(i)
	//				if err != nil {
	//					errChan <- err
	//				}
	//				select {
	//				case <-done:
	//					return
	//				case resStream <- i:
	//				}
	//			}(i)
	//		}
	//	}()
	//
	//	return resStream
	//}

	redisWork := func(done <-chan interface{}, errChan chan error, info <-chan int) <- chan interface{} {
		resStream := make(chan interface{})
		go func() {
			defer close(resStream)
			for i := range info {
				err := getInfo(i)
				if err != nil {
					errChan <- err
					resStream <- i
					continue
				}
				select {
				case <-done:
					return
				case resStream <- i:
				}
			}
		}()
		return resStream
	}

	mySqlFanIn := func(done <-chan interface{}, errChan chan error, limit int, channels []<-chan int) <-chan int {
		var wg sync.WaitGroup
		multiplexedResStream := make(chan int)

		multiplex := func(c <-chan int) {
			defer wg.Done()
			for i := range c {
				fmt.Println(i, " start")
				err := getInfo(i)
				if err != nil {
					errChan <- err
				}
				select {
				case <-done:
					return
				case multiplexedResStream <- i:
				}
			}
		}

		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}

		go func() {
			wg.Wait()
			close(multiplexedResStream)
		}()

		return multiplexedResStream
	}

	done := make(chan interface{})
	defer close(done)
	errChan := make(chan error)
	defer close(errChan)

	max := 1000
	limit := 100
	maxGoroutine := 10
	count := max / limit + 1
	lenStream := 0
	if count > maxGoroutine {
		lenStream = maxGoroutine
	} else {
		lenStream = count
	}

	offsetStream := make([]<-chan int, lenStream)
	chanMax := max / lenStream
	startOffset := 0
	for i := 0; i < lenStream; i++ {
		startOffset = i * chanMax
		offsetStream[i] = generator(done, startOffset, chanMax, limit)
	}
	pipeline := redisWork(done, errChan, mySqlFanIn(done, errChan, limit, offsetStream))

	for v := range pipeline {
		select {
		case err := <-errChan:
			fmt.Println("error:", err)
			done <- "done"
			return
		default:
			// redis
			fmt.Println(v)
		}
	}
	fmt.Printf("Search took: %v", time.Since(start))
}

func t2() {
	all := 1000
	from := 0

	//done := make(chan struct{})
	errChan := make(chan error)
	results := make(chan int, 11)
	working := make(chan struct{}, 11)
	//defer close(done)
	defer close(errChan)
	defer close(working)
	defer close(results)
	num := 100

	for from < all && len(working) <= 10 {
		from += num
		go worker(from, num, results, working, errChan)
	}
	for {
		select {
		//case <-done:
		//	fmt.Println("done!")
		//	return
		case <-errChan:
			fmt.Println("error!")
			break
		case <-results:
			if from < all {
				from += num
				go worker(from, num, results, working, errChan)
			}
			if from >= all && len(working) == 0 {
				fmt.Println("done!")
				return
			}
		}
	}

}

func worker(from, num int, results chan int, working chan struct{}, errChan chan error) {
	working <- struct{}{}
	fmt.Println("from: ", from, "num: ", num)
	time.Sleep(3 * time.Second)
	fmt.Println("from: ", from, "done")
	results <- from
	<- working
}
