package main

import (
	goquery2 "awesomeProject/github.com/PuerkitoBio/goquery"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	/*newurl := "https://search.bilibili.com/all?keyword=%E5%8D%8E%E4%B8%AD%E5%B8%88%E8%8C%83%E5%A4%A7%E5%AD%A6&from_source=banner_search"
	req, err := http.NewRequest("GET",newurl,nil)
	req.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
	rep, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	/*content, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		panic(err)
	}*/
	//defer rep.Body.Close()
	//fmt.Println(string(content))
	/*dom, err := goquery2.NewDocumentFromReader(rep.Body)

	*/
	for i := 1 ; i < 5 ; i++ {
		fmt.Printf("第 %d 页：\n",i)
		newurl := "https://search.bilibili.com/all?keyword=%E5%8D%8E%E4%B8%AD%E5%B8%88%E8%8C%83%E5%A4%A7%E5%AD%A6&from_source=banner_search"+"&page="+strconv.Itoa(i)
		req, err := http.NewRequest("GET",newurl,nil)
		req.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
		rep, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		dom, err := goquery2.NewDocumentFromReader(rep.Body)
		dom.Find(".title").Each(func(i int, selection *goquery2.Selection) {
			fmt.Println(selection.Text())
		})
		defer rep.Body.Close()
	}

}

