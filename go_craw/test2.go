package main

import (
	"awesomeProject/github.com/goquery"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	newurl := "https://accounts.douban.com/j/mobile/login/basic"
	data := url.Values{}
	data.Set("name", "13167302162")
	data.Set("password", "ccnu200181dx")
	data.Set("remember", "false")
	data.Set("ck", "")
	data.Set("ticket", "")
	load := strings.NewReader(data.Encode())
	req, err := http.NewRequest("POST", newurl, load)
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
	if err != nil {
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
	cookie := res.Cookies()
	newurl = "https://www.douban.com/people/206516033/"
	req1, err := http.NewRequest("GET", newurl, nil)
	if err != nil {
		panic(err)
	}
	req1.Header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36")
	for _, cookieq := range cookie {
		req1.AddCookie(cookieq)
	}
	res1, err := http.DefaultClient.Do(req1)
	if err != nil {
		panic(err)
	}
	/*content,err := ioutil.ReadAll(res1.Body)
	if err != nil {
		panic(err)
	}*/
	dom, err := goquery.NewDocumentFromReader(res1.Body)
	dom.Find("a").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection.Text())
	})
	//fmt.Printf("%s",content)
}
