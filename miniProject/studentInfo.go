package main

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"

	//"github.com/PuerkitoBio/goquery"
	"github.com/asynccnu/data_service/pkg/errno"
	"golang.org/x/net/publicsuffix"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type AccountReqeustParams struct {
	lt         string
	execution  string
	_eventId   string
	submit     string
	JSESSIONID string
}

type Info struct {
	Items []*InfoItem `json:"items" binding:"required"`
}

type InfoItem struct {
	xslb string `json:"kcmc" binding:"required"` // 学生类别(本科/研究)
	jgmc string `json:"kcmc" binding:"required"` // 学院名称
}

func main() {
	params,err := MakeAccountPreflightRequest()
	if err != nil {
		log.Println(err)
	}

	jar,err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Println(err)
	}
	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
		Jar:     jar,
	}
	//fmt.Println(params)
	err = MakeAccountRequest( "", "", params, &client)
	//err := MakeAccountRequest( "", "", params, &client)
	log.Println(err)
	//MakeXKRequest(&client)
	pt, err := MakeONERequest(&client)
	//fmt.Println(pt)
	pt = "Bearer " + pt
	fmt.Println(pt)

	GetInfo(pt)
	//MakeUndergraduateInfoRequest(&client,"2020","12")
	//a, _ := MakeInfoRequest(&client,"2019","3")

	//a, _ :=MakeUndergraduateInfoRequest(&client,"2020","3")
	//fmt.Println(a)
}

func MakeONERequest(client *http.Client) (portal_token string, err error) {
	request, err := http.NewRequest("GET", "http://one.ccnu.edu.cn", nil)
	if err != nil {
		log.Println(err)
		return "",err
	}

	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return "",err
	}

	u, err := url.Parse("http://one.ccnu.edu.cn")
	if err != nil {
		log.Println(err)
		return "",err
	}

	var pt string

	for _, cookie := range client.Jar.Cookies(u) {
		if cookie.Name == "PORTAL_TOKEN" {
			pt = cookie.Value
		}
		fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
	}
	return pt,nil
}

// 预处理，打开 account.ccnu.edu.cn 获取模拟登陆需要的表单字段
func MakeAccountPreflightRequest() (*AccountReqeustParams, error) {
	var JSESSIONID string
	var lt string
	var execution string
	var _eventId string

	params := &AccountReqeustParams{}

	// 初始化 http client
	client := http.Client{
		//Timeout: TIMEOUT,
	}

	// 初始化 http request
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login", nil)
	if err != nil {
		log.Println(err)
		return params, err
	}

	// 发起请求
	resp, err := client.Do(request)
	if err != nil {

		log.Println(err)
		return params, err
	}

	// 读取 Body
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Println(err)
		return params, err
	}

	// 获取 Cookie 中的 JSESSIONID
	for _, cookie := range resp.Cookies() {
		fmt.Println(cookie.Value)
		if cookie.Name == "JSESSIONID" {
			JSESSIONID = cookie.Value
		}
	}

	if JSESSIONID == "" {
		log.Println("Can not get JSESSIONID")
		return params, errors.New("Can not get JSESSIONID")
	}

	// 正则匹配 HTML 返回的表单字段
	ltReg := regexp.MustCompile("name=\"lt\".+value=\"(.+)\"")
	executionReg := regexp.MustCompile("name=\"execution\".+value=\"(.+)\"")
	_eventIdReg := regexp.MustCompile("name=\"_eventId\".+value=\"(.+)\"")

	bodyStr := string(body)

	ltArr := ltReg.FindStringSubmatch(bodyStr)
	if len(ltArr) != 2 {
		log.Println("Can not get form paramater: lt")
		return params, errors.New("Can not get form paramater: lt")
	}
	lt = ltArr[1]

	execArr := executionReg.FindStringSubmatch(bodyStr)
	if len(execArr) != 2 {
		log.Println("Can not get form paramater: execution")
		return params, errors.New("Can not get form paramater: execution")
	}
	execution = execArr[1]

	_eventIdArr := _eventIdReg.FindStringSubmatch(bodyStr)
	if len(_eventIdArr) != 2 {
		log.Println("Can not get form paramater: _eventId")
		return params, errors.New("Can not get form paramater: _eventId")
	}
	_eventId = _eventIdArr[1]

	log.Println("Get params successfully", lt, execution, _eventId)

	params.lt = lt
	params.execution = execution
	params._eventId = _eventId
	params.submit = "LOGIN"
	params.JSESSIONID = JSESSIONID

	return params, nil
}

func MakeAccountRequest(sid, password string, params *AccountReqeustParams, client *http.Client) error {
	v := url.Values{}
	v.Set("username", sid)
	v.Set("password", password)
	v.Set("lt", params.lt)
	v.Set("execution", params.execution)
	v.Set("_eventId", params._eventId)
	v.Set("submit", params.submit)
	//fmt.Println(strings.NewReader(v.Encode()))
		request, err := http.NewRequest("POST", "https://account.ccnu.edu.cn/cas/login;jsessionid="+params.JSESSIONID, strings.NewReader(v.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")


	resp, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return err
	}

	// check
	reg := regexp.MustCompile("class=\"success\"")
	matched := reg.MatchString(string(body))
	if !matched {
		log.Println("Wrong sid or pwd")
		return errno.ErrAuthFailed
	}

	log.Println("Login successfully")
	return nil
}

func MakeInfoRequest(client *http.Client, xnm, xqm string) (*Info, error) {
	var xslb string
	var jgmc string

	v := url.Values{}

	v.Set("xnm", xnm)
	v.Set("xqm", xqm)
	v.Set("_search", "false")
	v.Set("nd", string(time.Now().UnixNano()))
	v.Set("queryModel.showCount", "200")
	v.Set("queryModel.currentPage", "1")
	v.Set("queryModel.sortName", "")
	v.Set("queryModel.sortOrder", "asc")
	v.Set("time", "0")

	request, err := http.NewRequest("POST", "http://one.ccnu.edu.cn/index#/app/home/user_center", strings.NewReader(v.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=UTF-8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("Origin", "http://xk.ccnu.edu.cn")
	request.Header.Set("Host", "xk.ccnu.edu.cn")
	request.Header.Set("Referer", "http://one.ccnu.edu.cn/index")

	resp, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}
	content,_ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(content))
	html, err := goquery.NewDocumentFromReader(strings.NewReader(string(content)))
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	html.Find(".user_center-cont").Find(".col-xs-9").Find(".row").Find("div").Each(func(i int, selection *goquery.Selection) {
		if i == 7 {
			jgmc = string(selection.Text())
		}
		if i == 13 {
			xslb = string(selection.Text())
		}
	})

	var info = &Info{}

	fmt.Println(xslb);fmt.Println(jgmc)
	fmt.Println(info)
	if len(info.Items) == 0 {
		return nil, errors.New("empty info list")
	}
	return info, nil
}

func GetInfo(pt string) {
	client1 := http.Client{}
	v := url.Values{}
	request, _ := http.NewRequest("POST","http://one.ccnu.edu.cn/user_portal/index",strings.NewReader(v.Encode()))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.Header.Set("Authorization", pt)
	resp, err := client1.Do(request)
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