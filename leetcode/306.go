package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isAdditiveNumber("101"))
}

func isAdditiveNumber(num string) bool {
	var now []int
	var res bool

	verify := func() bool {
		if len(now) < 3 {
			return true
		}
		if now[len(now)-1] == now[len(now)-2] + now[len(now)-3] {
			return true
		}
		return false
	}
	upperBound := func() int {
		if len(now) < 2 {
			return math.MaxInt64
		}
		return now[len(now)-1] + now[len(now)-2]
	}

	var backtrace func(index int)
	backtrace = func(index int) {
		if res {
			return
		}
		if index == len(num) {
			if len(now) >= 3 {
				res = true
				fmt.Println(now)
			}
			return
		}
		if num[index] == '0' {
			now = append(now, 0)
			if verify() {
				backtrace(index+1)
			}
			now = now[:len(now)-1]
			return
		}

		nowNum := 0
		ub := upperBound()
		for i := index; i < len(num) && nowNum <= ub ; i ++ {
			nowNum = nowNum*10 + int(num[i]-'0')
			now = append(now, nowNum)
			if verify() {
				backtrace(i+1)
			}
			now = now[:len(now)-1]
		}
	}

	backtrace(0)

	return res
}
