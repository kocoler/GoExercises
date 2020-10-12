package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	num := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000)
	fmt.Println(num)
	code := strconv.FormatInt(int64(num), 10)
	fmt.Println(code)
}
