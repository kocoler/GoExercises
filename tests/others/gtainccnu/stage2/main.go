package main

import (
	"fmt"

	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/encrypt"
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/secret_key", "", httptool.DEFAULT)
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

	extrainfo := resp.Body.Data.ExtraInfo
	d, err := encrypt.Base64Decode(extrainfo)
	if err != nil {
		panic(err)
	}
	fmt.Println(d)
	e, err := encrypt.AESEncryptOutInBase64([]byte("for {go func(){time.Sleep(1*time.Hour)}()}"), []byte("MuxiStudio203304"))
	if err != nil {
		panic(err)
	}
	req, err = httptool.NewRequest(httptool.PUTMETHOD, "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/bank/gate", string(e), httptool.DEFAULT)
	if err != nil {
		panic(err)
	}
	req.AddHeader("passport", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjb2RlIjoibWFjIiwiaWF0IjoxNjMzNDAyMzY2LCJuYmYiOjE2MzM0MDIzNjZ9.lfJitrB-jX8_vWuzhY4NZVYzASwJwtfjATOZs2TW0P4")
	resp, err = httptool.SendRequest(req)
	if err != nil {
		panic(err)
	}
	resp.ShowBody()
	resp.ShowHeader()
}
