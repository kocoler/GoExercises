package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	reuqsrURL := "https://www.baidu.com/s?ie=utf-8&f=8&rsv_bp=1&tn=ubuntuu_cb&wd=go%E6%AD%A3%E5%88%99%E8%A1%A8%E8%BE%BE%E5%BC%8F&oq=%25E8%25B1%2586%25E7%2593%25A3%25E8%25AF%25BB%25E4%25B9%25A6&rsv_pq=916cde7300454470&rsv_t=6ebe97tcPXdLLOzf92Kl2naAh2nV6LrMXMM7%2BV3aBaeunZqjzb%2FonePuOUD3RApyOg&rqlang=cn&rsv_enter=1&rsv_dl=tb&inputT=10890&rsv_sug3=32&rsv_sug1=22&rsv_sug7=101&rsv_sug2=0&rsv_sug4=14721"
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
