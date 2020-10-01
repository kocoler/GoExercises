package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

func main() {
	//2020-07-21T11:07:49+08:00
	fmt.Println(strconv.FormatInt(time.Now().Unix(), 10))
	t, err := time.Parse("2006-01-02T15:04:05+08:00", "2020-07-26T16:49:37+08:00")
	h, _ := time.ParseDuration("-1h")
	t = t.Add(8 * h)
	now := time.Now()
	log.Println(now.Sub(t).String(), now.String(),t.String(), err)
}
