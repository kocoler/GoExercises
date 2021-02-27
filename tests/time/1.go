package main

import (
	"fmt"
	"time"
)

func main() {
	t, err := time.Parse(time.RFC850, time.RFC850)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t.Unix())
}

