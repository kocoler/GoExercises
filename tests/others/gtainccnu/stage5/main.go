package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest("", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination", "./submissions.go", httptool.FILE)
	if err != nil {
		panic(err)
	}

	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoibWFjIiwiaWF0IjoxNjMzNDAyMzY2LCJuYmYiOjE2MzM0MDIzNjZ9.lfJitrB-jX8_vWuzhY4NZVYzASwJwtfjATOZs2TW0P4")
	//resp, err := httptool.SendRequest(req)
	//if err != nil {
	//	panic(err)
	//}
	//resp.ShowBody()
	//resp.ShowHeader()
	client := &http.Client{}
	response, err := client.Do(req.Req)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(string(body))
}
