package main

import (
	"encoding/json"
	"errors"
	"fmt"
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

var TIMEOUT = time.Duration(30 * time.Second)

// Grade ... 成绩数据结构
type Grade struct {
	Items []GradeItem `json:"items" binding:"required"`
}

// GradeItem ... 成绩数据项
type GradeItem struct {
	Kcmc   string `json:"kcmc" binding:"required"`   // 课程名称
	Kcxzmc string `json:"kcxzmc" binding:"required"` // 课程性质名称，比如专业主干课程/通识必修课
	Cj     string `json:"cj" binding:"required"`     // 成绩
	Xf     string `json:"xf" binding:"required"`     // 学分
	Jsxm   string `json:"jsxm" binding:"required"`   // 教师姓名
	Kclbmc string `json:"kclbmc" binding:"required"` // 课程类别名称，比如专业课/公共课
	Kcgsmc string `json:"kcgsmc" binding:"required"` // 课程归属名称，比如文/理
	//KchID  string `json:"kch_id"`                    // 课程号id
	JxbID  string `json:"jxb_id"`                    // 教学班id
	Zymc   string `json:"zymc"`                      // 学生专业名称
	Jgmc   string `json:"jgmc"`                      // 学生学院名称
	//JxbId  string `json:"jxb_id" binding:"required"`
	Xnm    string `json:"xnm" binding:"required"`    // 学年名称，比如 2016（可能返回的课程列表包含多个学年）
}

type AccountReqeustParams struct {
	lt         string
	execution  string
	_eventId   string
	submit     string
	JSESSIONID string
}

// 获取成绩请求
type GradeRequest struct {
	Sid                  string   `protobuf:"bytes,1,opt,name=sid,proto3" json:"sid,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Xqm                  string   `protobuf:"bytes,2,opt,name=xqm,proto3" json:"xqm,omitempty"`
	Xnm                  string   `protobuf:"bytes,3,opt,name=xnm,proto3" json:"xnm,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type classes struct {
	Classname string
	JxbId     string
}

type ResultDetails struct {
	ClassName  string
	DailyScoreRate string
	DailyScore string
	FinalScoreRate string
	FinalScore string
	TotalScore string
}

func main() {
	sid := ""
	pwd := ""
	xqm := ""
	xnm := ""
	cookies, class := GetCookiesAndClassname(sid, pwd, xqm, xnm)
	//fmt.Println(cookies,class)
	var result []ResultDetails
	for _,v := range class {
		//fmt.Println(v)
		tmpStr := GetDetails(cookies, v.Classname, xnm, xqm, v.JxbId)
		tmpResult := ProcessExp(tmpStr)
		tmpResult.ClassName = v.Classname
		result = append(result, tmpResult)
	}
	fmt.Println(result)
}

//获取详细的单科页面数据
func GetDetails(cookies string, classname string, xnm string, xqm string, jxbId string) string {
	data := url.Values{}
	data.Set("xnm",xnm)
	data.Set("xqm",xqm)
	data.Set("kcmc",classname)
	data.Set("jxb_id",jxbId)

	//newurl := "http://xk.ccnu.edu.cn/cjcx/cjcx_cxCjxq.html?"+string(time.Now().Unix())+"&gnmkdm=N305005"
	req, err := http.NewRequest("POST","http://xk.ccnu.edu.cn/cjcx/cjcx_cxCjxq.html?time=1579345703&gnmkdm=N305005",strings.NewReader(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-gorm-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	req.Header.Set("Cookie","JSESSIONID="+cookies)
	//req.Header.Set("Host","xk.ccnu.edu.cn")
	//req.Header.Set("Connection","keep-alive")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(body))
	return string(body)
}

//进行正则分析
func ProcessExp(body string) ResultDetails {
	reg1 := regexp.MustCompile("<td valign=\"middle\">(.*?)</td>")
	result := reg1.FindAllStringSubmatch(body,-1)
	var tmpResult ResultDetails
	i := 0
	for _, text := range result {
		if i == 1 {
			tmpResult.DailyScoreRate = ChangeString(text[1], 1, len(text[1])-5)
			i = 2
			continue
		}
		if i == 2 {
			tmpResult.DailyScore = ChangeString(text[1], 1, len(text[1])-5)
			i = 0
		}
		if i == 3 {
			//fmt.Println(text[1],len(text[1]),ChangeString(text[1], 1, len(text[1])-6) )
			tmpResult.FinalScoreRate = ChangeString(text[1], 1, len(text[1])-5)
			i = 4
			continue
		}
		if i == 4 {
			tmpResult.FinalScore = ChangeString(text[1], 1, len(text[1])-5)
			i = 0
		}
		if i == 5 {
			i ++
			continue
		}
		if i == 6 {
			tmpResult.TotalScore = ChangeString(text[1], 1, len(text[1])-5)
			i = 0
		}
		if text[1] == "【 平时 】" {
			i = 1
		}
		if text[1] == "【 期末 】" {
			i = 3
		}
		if text[1] == "【 总评 】" {
			i = 5
		}
	}
	return tmpResult
}

func GetCookiesAndClassname(sid, pwd, xqm, xnm string) (string, []classes) {
	var result []classes
	params, err := MakeAccountPreflightRequest()
	if err != nil {
		log.Println(err)
	}

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Println(err)
	}

	client := http.Client{
		Timeout: time.Duration(10 * time.Second),
		Jar:     jar,
	}

	in2 := GradeRequest{
		Sid:                  sid,
		Password:             pwd,
		Xqm:                  xqm,
		Xnm:                  xnm,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}

	in := &in2

	if err := MakeAccountRequest(in.Sid, in.Password, params, &client); err != nil {

	}

	var cookies string
	if cookies, err = MakeXKRequest(&client); err != nil {

	}

	if grade, err := MakeGradeRequest(&client, in.Xnm, in.Xqm); err != nil {
		MakeAccountLogoutRequest(&client)
		log.Println("logout success")

	} else {
	//	MakeAccountLogoutRequest(&client)
		log.Println("logout success")
		//lists := make([]pb.GradeItem, 0)
		for i := 0; i < len(grade.Items); i++ {
		/*	item := grade.Items[i]
			lists = append(lists, pb.GradeItem{
				Kclbmc: item.Kclbmc,
				Kcmc:   item.Kcmc,
				Cj:     item.Cj,
				Jsxm:   item.Jsxm,
				Kcxzmc: item.Kcxzmc,
				Xf:     item.Xf,
				Kcgsmc: item.Kcgsmc,
				Xnm:    item.Xnm,

			})
		 */
			item := grade.Items[i]
			tmpClass := classes{
				Classname: item.Kcmc,
				JxbId:     item.JxbID,
			}
			result = append(result, tmpClass)
		}

		//GetDetails(cookies, "什么是科学", &client)
		return cookies, result
		//return &pb.GradeReply{Lists: lists}, nil
	}
	return cookies, result
}

// 退出信息门户登录
func MakeAccountLogoutRequest(client *http.Client) error {
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/logout", nil)
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// account.ccnu.edu.cn 模拟登录，用于验证账号密码是否可以正常登录
func MakeAccountRequest(sid, password string, params *AccountReqeustParams, client *http.Client) error {
	v := url.Values{}
	v.Set("username", sid)
	v.Set("password", password)
	v.Set("lt", params.lt)
	v.Set("execution", params.execution)
	v.Set("_eventId", params._eventId)
	v.Set("submit", params.submit)

	request, err := http.NewRequest("POST", "https://account.ccnu.edu.cn/cas/login;jsessionid="+params.JSESSIONID, strings.NewReader(v.Encode()))
	fmt.Println(strings.NewReader(v.Encode())) //
	request.Header.Set("Content-Type", "application/x-www-gorm-urlencoded")
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

// 预处理，打开 account.ccnu.edu.cn 获取模拟登陆需要的表单字段
func MakeAccountPreflightRequest() (*AccountReqeustParams, error) {
	var JSESSIONID string
	var lt string
	var execution string
	var _eventId string

	params := &AccountReqeustParams{}

	// 初始化 http client
	client := http.Client{
		Timeout: TIMEOUT,
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
		log.Println("Can not get gorm paramater: lt")
		return params, errors.New("Can not get gorm paramater: lt")
	}
	lt = ltArr[1]

	execArr := executionReg.FindStringSubmatch(bodyStr)
	if len(execArr) != 2 {
		log.Println("Can not get gorm paramater: execution")
		return params, errors.New("Can not get gorm paramater: execution")
	}
	execution = execArr[1]

	_eventIdArr := _eventIdReg.FindStringSubmatch(bodyStr)
	if len(_eventIdArr) != 2 {
		log.Println("Can not get gorm paramater: _eventId")
		return params, errors.New("Can not get gorm paramater: _eventId")
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

// xk.ccnu.edu.cn 模拟登录，用于请求成绩/课表等等
func MakeXKRequest(client *http.Client) (string, error) {
	var result string
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login?service=http%3A%2F%2Fxk.ccnu.edu.cn%2Fssoserver%2Flogin%3Fywxt%3Djw%26url%3Dxtgl%2Findex_initMenu.html", nil)
	if err != nil {
		log.Println(err)
		return result, err
	}

	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return result, err
	}

	u, err := url.Parse("http://xk.ccnu.edu.cn")
	if err != nil {
		log.Println(err)
		return result, err
	}

	for _, cookie := range client.Jar.Cookies(u) {
		//fmt.Printf("  %s: %s\n", cookie.Name, cookie.Value)
		return cookie.Value, nil
	}
	return result, nil
}

// MakeGradeRequest ... xk.ccnu.edu.cn 获取成绩
func MakeGradeRequest(client *http.Client, xnm, xqm string) (Grade, error) {
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

	request, err := http.NewRequest("POST", "http://xk.ccnu.edu.cn/cjcx/cjcx_cxDgXscj.html?doType=query&gnmkdm=N305005", strings.NewReader(v.Encode()))
	request.Header.Set("Content-Type", "application/x-www-gorm-urlencoded;charset=UTF-8")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.109 Safari/537.36")
	request.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	request.Header.Set("X-Requested-With", "XMLHttpRequest")
	request.Header.Set("Origin", "http://xk.ccnu.edu.cn")
	request.Header.Set("Host", "xk.ccnu.edu.cn")
	request.Header.Set("Referer", "http://xk.ccnu.edu.cn//cjcx/cjcx_cxDgXscj.html?gnmkdm=N305005&layout=default&su=2016210942")

	resp, err := client.Do(request)
	if err != nil {
		log.Print(err)
	}
	var grade = Grade{}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Println(err)
		return grade, err
	}

	//fmt.Println(string(body))

	if err := json.Unmarshal(body, &grade); err != nil {
		return grade, err
	}
	//fmt.Println(grade)
	if len(grade.Items) == 0 {
		return grade, errors.New("empty grade list")
	}

	return grade, nil
}

//截取字符串长度小函数
func ChangeString(str string, from int, end int) string {
	if len(str) == 0 || end < from {
		return ""
	}
	tmpStr := []byte(str)
	tmpStr1 := tmpStr[from-1 : end-1]
	return string(tmpStr1)
}