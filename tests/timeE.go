package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// 1612252334
	//fmt.Println(time.Now().Add(1 * time.Second).Unix())
	//fmt.Println(time.Now().Unix() - 1612252364)
	fmt.Println(time.Unix(1136214245, 0))
	t, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:04:05")
	fmt.Println(t.Unix())
}

func t1() {
	timestr := "2020-05-02T01:00:00+0800"
	//var n string
	timestr = strings.ReplaceAll(timestr, "T", " ")
	times, _ := time.Parse("2006-01-02 15:04:05+0000", timestr)
	fmt.Println(timestr[:19])
	timeUnix := times.Unix()
	fmt.Printf("times is %+v, timeUnix is %+v", times, timeUnix)
}
