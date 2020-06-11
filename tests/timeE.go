package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	timestr := "2020-05-02T01:00:00+0800"
	//var n string
	timestr = strings.ReplaceAll(timestr, "T", " ")
	times, _ := time.Parse("2006-01-02 15:04:05+0000", timestr)
	fmt.Println(timestr[:19])
	timeUnix := times.Unix()
	fmt.Printf("times is %+v, timeUnix is %+v", times, timeUnix)
}
