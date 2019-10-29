package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	reuqsrURL := "https://baike.baidu.com/item/%E7%BD%91%E7%BB%9C%E7%88%AC%E8%99%AB/5162711?fromtitle=%E7%88%AC%E8%99%AB&fromid=22046949&fr=aladdin"
	re, err := http.Get(reuqsrURL)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(re.Body)
	if err != nil {
		panic(err)
	}
	content := string(body)
	defer re.Body.Close()
	fmt.Println(content)
}
