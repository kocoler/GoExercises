package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	client := http.Client{}
	v := url.Values{}
	request, _ := http.NewRequest("POST","http://one.ccnu.edu.cn/user_portal/index",strings.NewReader(v.Encode()))
	request.Header.Set("Content-Type", "application/x-www-gorm-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiI4RUNFRkEyQjBFNkI2MjFDRTA1MDAwMDAwMDAwNDlEMiIsImV4cCI6MTU3OTM0NTc2Mn0.F1QIC45FY4-Vxm01FDCqrvtvq_4VgIzu0ek9xtIgxag")
	resp, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(body))
}


