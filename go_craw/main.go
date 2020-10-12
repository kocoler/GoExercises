package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "https://book.douban.com/"
	//re, err := http.Get(reuqsrURL)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//panic(err)
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.87 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		//panic(err)
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(body)

	/*if err != nil {
		panic(err)
	}*/

	//defer re.Body.Close()

	//reg := regexp.MustCompile(`[\p{Han}]+`)   // 查找连续的汉字
	//reg := regexp.MustCompile(`title=[...]<`)
	//fmt.Println( reg.FindAllStringSubmatch(content, -1))
	//fmt.Println(content)
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Fatalln(err)
	}
	dom.Find(".name").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
}
