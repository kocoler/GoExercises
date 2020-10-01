package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var A string = "https://graph.qq.com/oauth2.0/token"

type qqToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type qqOpenId struct {
	ClientId string `json:"client_id"`
	OpenId  string `json:"open_id"`
}

//FE7A928CD9AF1A873874566E1148454A
func main() {
	/*
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	AppId := "101896710"
	params.Add("client_id", AppId)
	AppKey := "8d3d6df6f0f5765977a819de358ca788"
	params.Add("client_secret", AppKey)
	code := "52052E0B7EF7C85508BCD4A6E4B8C81B"
	params.Add("code", code)
	redirectURI := "https://forum.casbin.com/callback/qq/signup"
	str := fmt.Sprintf("%s&redirect_uri=%s", params.Encode(), redirectURI)
	loginURL := fmt.Sprintf("%s?%s", "https://graph.qq.com/oauth2.0/token", str)

	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, loginURL, nil); err != nil {
		log.Println(err)
		//return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		log.Println(err)
		//return nil, err
	}

	// 将响应体解析为 token，并返回
	var token qqToken
	if err = json.NewDecoder(res.Body).Decode(&token); err != nil {
		log.Println(err)
		//return nil, err
	}
*/
	/*
	accessToken := "DEB03181FDEADF1BD015562F683852C3"
	getOpenIdUrl := fmt.Sprintf(
		"https://graph.qq.com/oauth2.0/me?access_token=%s", accessToken,
	)
	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, getOpenIdUrl, nil); err != nil {
		log.Println(err)
		//return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		log.Println(err)
		//return nil, err
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(contents))
	// 将响应体解析为 token，并返回
	reg1 := regexp.MustCompile("\"openid\":\"(.*?)\"}")
	result1 := reg1.FindAllStringSubmatch(string(contents), -1)


	fmt.Println(result1)

	 */


	accessToken := "DEB03181FDEADF1BD015562F683852C3"
	openId := "171879FF24A9D54E482D040CE35AE57D"
	AppId := "101896710"

	getUserInfo := fmt.Sprintf(
		"https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s", accessToken, AppId, openId,
	)

	var req *http.Request
	var err error
	if req, err = http.NewRequest(http.MethodGet, getUserInfo, nil); err != nil {
		log.Println(err)
		//return nil, err
	}
	req.Header.Set("accept", "application/json")

	// 发送请求并获得响应
	var httpClient = http.Client{}
	var res *http.Response
	if res, err = httpClient.Do(req); err != nil {
		log.Println(err)
		//return nil, err
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(contents))
	// 将响应体解析为 token，并返回
	//reg1 := regexp.MustCompile("\"openid\":\"(.*?)\"}")
	//result1 := reg1.FindAllStringSubmatch(string(contents), -1)


	//err = json.Unmarshal(contents, &)
	//if err != nil {
	//	log.Println(err)
	//}

	//fmt.Println(result1)

	//resultMap := convertToMap(body)

	//info := &PrivateInfo{}
	//info.AccessToken = resultMap["access_token"]
	//info.RefreshToken = resultMap["refresh_token"]
	//info.ExpiresIn = resultMap["expires_in"]
	//response, err := http.Get(A+?access_token=" + token.AccessToken)

	//defer response.Body.Close()
	//contents, err := ioutil.ReadAll(response.Body)
	//fmt.Fprintf(w, "Content: %s\n", contents)
}
