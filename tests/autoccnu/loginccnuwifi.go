package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

var s = flag.String("s", "", "Your sid.")
var p = flag.String("p", "", "Your password.")
var d = flag.Bool("d", false, "Whether output information.")

func outputWaring(info string) {
	fmt.Printf("%c[%d;%dm%s%c[0m  ", 0x1B, 34, 1, "[WARNING]", 0x1B)
	fmt.Printf("%c[%d;%dm%s%c[0m\n", 0x1B, 37, 1, info, 0x1B)
}

func outputErr(info string) {
	fmt.Printf("%c[%d;%dm%s%c[0m  ", 0x1B, 31, 1, "[ERR]", 0x1B)
	fmt.Printf("%c[%d;%dm%s%c[0m\n", 0x1B, 37, 1, info, 0x1B)
}

func outputInfo(info string) {
	fmt.Printf("%c[%d;%dm%s%c[0m  ", 0x1B, 35, 1, "[INFO]", 0x1B)
	fmt.Printf("%c[%d;%dm%s%c[0m\n", 0x1B, 37, 1, info, 0x1B)
}

func checkErr(str string) []string {
	reg := regexp.MustCompile("errmsg=(.*?)\"")
	result := reg.FindAllStringSubmatch(str, -1)

	var res []string
	for _, v := range result {
		info, _ := url.QueryUnescape(v[1])
		res = append(res, info)
	}

	return res
}

func main() {
	//for i, args := range os.Args {
	//	fmt.Printf("args[%d]=%s\n",i,args)
	//}

	// fmt.Println(len(os.Args))
	flag.Parse()

	if len(os.Args) < 5 {
		outputWaring("usage: autoccnu -s <yourSid> -p <yourPassword>")
		outputInfo("autoccnu -h for more information")
		return
	}

	err := login(*s, *p)
	if err != nil {
		info := checkErr(err.Error())
		for _, v := range info {
			outputErr(v)
		}
        if len(info) == 0 {
            outputErr(err.Error())
        }
		return
	}

	outputInfo("auto login ccnu wifi success!")
}

func login(sid, pwd string) error {
	url := "http://securelogin.arubanetworks.com/auth/index.html/u"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf("user=%s&password=%s", sid, pwd))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return err
	}
	req.Header.Add("Host", "securelogin.arubanetworks.com")
	req.Header.Add("Content-Length", "35")
	req.Header.Add("Cache-Control", "max-age=0")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("Origin", "http://login.ccnu.edu.cn")
	req.Header.Add("Content-Type", "application/x-www-gorm-urlencoded")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Safari/537.36")
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Referer", "http://login.ccnu.edu.cn/")
	// req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,en-US;q=0.6")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	// fmt.Println(string(body))

	return nil
}
