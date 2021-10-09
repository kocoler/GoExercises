package main

import (
	"time"
)

func main() {
	a := 0
	for i := 0; i < 100; i ++ {
		go func() {
			a ++
			time.Sleep(10 * time.Second)
		}()

		go func() {
			b := 0
			b ++
			time.Sleep(10 * time.Second)
		}()
	}

	time.Sleep(20 * time.Second)
}
