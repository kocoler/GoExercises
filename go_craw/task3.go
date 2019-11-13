package main

import (
	"awesomeProject/github.com/goquery"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func main() {
	newurl := "https://accounts.douban.com/j/mobile/login/basic"
	data := url.Values{}
	data.Set("name","13167302162")
	data.Set("password","ccnu200181dx")
	data.Set("remember","false")
	data.Set("ck","")
	data.Set("ticket","")
	load := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST",newurl,load)
	if err != nil {
		panic(err)
		return
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Origin", "https://accounts.douban.com/passport/login_popup?login_source=anony")
	req.Header.Add("Referer", "https://accounts.douban.com/passport/login_popup?login_source=anony")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	res, err := http.DefaultClient.Do(req)
	if err != nil  {
		panic(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
		return
	}
	fmt.Printf(string(body))
	/*reuqsrURL := "https://book.douban.com/top250?icn=index-book250-all"
	req1, err := http.NewRequest("GET",reuqsrURL,nil)
	req1.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
	//req1.Header.Add("Cookic","%228%2FRlDH9JEDk%22")
	/*if err != nil {
		panic(err)
	}*/
	/*rep1, err := http.DefaultClient.Do(req1)
	if err != nil {
		panic(err)
	}

	/*content, err := ioutil.ReadAll(rep1.Body)
	if err != nil {
		panic(err)
	}*/
	/*defer rep1.Body.Close()
	dom, err := goquery.NewDocumentFromReader(rep1.Body)
	/*dom.Find("tbody").Each(func(i int, selection *goquery.Selection) {
		//fmt.Println(selection.Text())
		dom.Find(".pl").Each(func(i int, selection2 *goquery.Selection) {
			//fmt.Println(selection.Text())
			//fmt.Println(selection2.Text())
		})

	})*/
	/*dom.Find(".pl").Each(func(i int, selection *goquery.Selection) {
		fmt.Printf("%s",selection.Text())
	})*/
	j := 1
	for i := 1; i <= 225; i+=24 {
		reuqsrURL := "https://book.douban.com/top250?start="+strconv.Itoa(i)
		req1, err := http.NewRequest("GET",reuqsrURL,nil)
		req1.Header.Add("User-Agent","Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
		//req1.Header.Add("Cookic","%228%2FRlDH9JEDk%22")
		/*if err != nil {
			panic(err)
		}*/
		rep1, err := http.DefaultClient.Do(req1)
		if err != nil {
			panic(err)
		}

		/*content, err := ioutil.ReadAll(rep1.Body)
		if err != nil {
			panic(err)
		}*/
		defer rep1.Body.Close()
		dom, err := goquery.NewDocumentFromReader(rep1.Body)
		/*dom.Find("tbody").Each(func(i int, selection *goquery.Selection) {
			//fmt.Println(selection.Text())
			dom.Find(".pl").Each(func(i int, selection2 *goquery.Selection) {
				//fmt.Println(selection.Text())
				//fmt.Println(selection2.Text())
			})

		})*/
		/*dom.Find(".pl").Each(func(i int, selection *goquery.Selection) {
			fmt.Printf("%s",selection.Text())
		})*/
		
		dom.Find("div[class=pl2]").Each(func(i int, selection3 *goquery.Selection) {
			//_,err := io.WriteString("task3.txt",selection3.Text())
			fmt.Printf("第 %d 位：",j)
			j ++
			fmt.Printf("%s",selection3.Text())
		})

	}
	//fmt.Println(string(content))
}

