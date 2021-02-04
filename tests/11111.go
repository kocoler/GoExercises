package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	time.Sleep(100 * time.Millisecond)
	var wg sync.WaitGroup

	s := 2017210001
	for i:=0; i < 80; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("请求%d\n", i)
			request(strconv.Itoa(s + i))
		}()
	}

	wg.Wait()
}

func request(sid string) {

	url := "http://47.97.74.180:8899/api/v1/data"
	method := "GET"

	payload := strings.NewReader(`{
    "id":"`+sid+`",
    "password":"MTU4MTE4NTIxMzM="
}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body[:100]))
	fmt.Println("done")
}
