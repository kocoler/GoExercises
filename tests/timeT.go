package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	//currentTime := time.Now()
	//res := currentTime.AddDate(0, -1, 0)
	//fmt.Println(res.Format(time.RFC3339))
	/*time1 := time.Date(2017,11,11,0,0,0,0,time.Local)
	fmt.Println(time1.Format("2006-01-02"))
	fmt.Println(time.Now())*/
	fmt.Println(GetTimeMinute(10))
}

func GetTimeMinute(minute int) string {
	currentTime := time.Now()
	m, _ := time.ParseDuration("-" + strconv.Itoa(minute) + "m")
	fmt.Println(m)
	res := currentTime.Add(m)
	return res.Format(time.RFC3339)
}
