package main

import (
	"github.com/Grand-Theft-Auto-In-CCNU-MUXI/hacker-support/httptool"
)

func main() {
	req, err := httptool.NewRequest("GET", "http://http-theft-bank.gtainccnu.muxixyz.com/api/v1/organization/code", "", httptool.DEFAULT)
	if err != nil {
		panic(err)
	}
	resp, err := req.SendRequest()
	if err != nil {
		panic(err)
	}
	resp.ShowBody()
	//resp.ShowHeader()
}
