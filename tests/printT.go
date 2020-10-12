package main

import (
	"fmt"
	"time"
)

func main() {
	tNow := time.Now()
	//时间转化为string，layout必须为 "2006-01-02 15:04:05"
	timeNow := tNow.Format("2006-01-02 15:04:05")
	fmt.Println("tNow(time format):", tNow)
	fmt.Println("tNow(string format):", timeNow)
	fmt.Printf("%T",timeNow)
}
