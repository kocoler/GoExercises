package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"unicode"
)

func main() {
	url := "https://ip.tool.lu"
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := strings.FieldsFunc(string(body), func(c rune) bool {
		if c == '.' {
			return false
		}

		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	fmt.Println(r)
}
