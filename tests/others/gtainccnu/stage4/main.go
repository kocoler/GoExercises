package main

import "github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"

func main() {
	req, err := httptool.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/muxi/backend/computer/examination", "", httptool.DEFAULT)
	if err != nil {
		panic(err)
	}
	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoibWFjIiwiaWF0IjoxNjMzNDAyMzY2LCJuYmYiOjE2MzM0MDIzNjZ9.lfJitrB-jX8_vWuzhY4NZVYzASwJwtfjATOZs2TW0P4")
	resp, err := httptool.SendRequest(req)
	if err != nil {
		panic(err)
	}
	resp.ShowBody()
	resp.ShowHeader()
}