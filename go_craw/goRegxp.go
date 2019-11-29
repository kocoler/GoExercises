package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	newurl := "https://search.bilibili.com/all?keyword=%E5%8D%8E%E4%B8%AD%E5%B8%88%E8%8C%83%E5%A4%A7%E5%AD%A6&from_source=banner_search"
	req, err := http.NewRequest("GET",newurl,nil)
	req.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
	rep, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(rep.Body)
	fmt.Println(string(content))
	exp1 := regexp.MustCompile(`</span><a title="(.*?)"\shref="//(.*?)"`)
	//exp2 := regexp.MustCompile(`class="up-name">(.*?)</a>`)
	//result2 := exp2.FindAllStringSubmatch(string(content),-1)
	result1 := exp1.FindAllStringSubmatch(string(content),-1)
	fmt.Println(result1)
	for i := 0; i < len(result1); i++ {
		//fmt.Println(result2[i][1])
		fmt.Println(result1[i][2])
	}
}

//title="【华中师范大学】最大蹦迪现场2016校运会开幕式DJ完整版" target="_blank"