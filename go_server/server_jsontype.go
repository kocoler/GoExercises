package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type tycontent struct {
	Name string `json:"name"`
	Password string `json:"-"`
	Gender string `json:"gender"`
	Favourite string `json:"favourite"`
}

var (
	data = make(map[string]*tycontent)
	i int
)

func main_body(w http.ResponseWriter, r *http.Request) {
	//r.FormValue()
	defer r.Body.Close()
	//http.Redirect(w,r,"http://localhost:9090/test",301)
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var titles struct {
		Name     string `json:"name"`
		Password string `json:"password,omitempty"`
	}
	err = json.Unmarshal(content, &titles)
	if err != nil {
		log.Println(err)
		//fmt.Fprintln(w,err)
		//return
	}
	user_cookie := &http.Cookie {
		Name:    "my_account",
		Value:   "111",
		Path: "/",
		Expires: time.Time{},
		MaxAge:  60,
	}

	//a,_ := r.Cookie("Name")
	//fmt.Fprintf(w,user_cookie.String())
	//fmt.Fprintf(w,"success")
	if data[titles.Name] == nil {
		password := titles.Password
		i++
		user_cookie.Value = titles.Name
		w.Header().Add("Set-Cookie",user_cookie.String())
		fmt.Fprintf(w, "注册用户\n")
		//fmt.Fprintln(w,password)
		data[titles.Name] = new(tycontent)
		data[titles.Name].Name = titles.Name
		data[titles.Name].Password = password
		fmt.Fprintf(w, "注册成功\n")
		fmt.Fprintf(w, "\n在主页面的第"+strconv.Itoa(i)+"次操作\n")
	} else {
		i++
		password := titles.Password
		if data[titles.Name].Password == password {
			user_cookie.Value = titles.Name
			w.Header().Add("Set-Cookie",user_cookie.String())
			fmt.Fprintf(w, "登入成功！\n")
			fmt.Fprintf(w, "\n在主页面的第"+strconv.Itoa(i)+"次操作\n")
		} else {
			user_cookie.Value = titles.Name
			w.Header().Add("Set-Cookie",user_cookie.String())
			fmt.Fprintf(w, "修改密码\n")
			data[titles.Name].Password = password
			fmt.Fprintf(w, "修改成功\n")
			fmt.Fprintf(w, "\n在主页面的第"+strconv.Itoa(i)+"次操作\n")
		}
	}
		//json.Unmarshal(content,&titles)
		//fmt.Fprintln(w,titles)
}
/*
func Content(v http.ResponseWriter, e *http.Request)  {

}
*/

func m_content(w1 http.ResponseWriter, r1 *http.Request)  {
	c := r1.Cookies()
	//c,_ := r.Cookie(titles.Name)
	if len(c) == 0 {
		fmt.Fprintln(w1,"请登陆")
		fmt.Fprintln(w1,"跳转中...")
		http.Redirect(w1,r1,"http://localhost:9090/body",301)
	}
	var cookiename string
	for _,cookieq := range c {
		cookiename = cookieq.Value
	}
	if len(data[cookiename].Name) == 0 {
		w1.WriteHeader(403)
		return
	}
	fmt.Fprintln(w1,"修改信息：\n")
	time.Sleep(100 *time.Microsecond)
	content, err := ioutil.ReadAll(r1.Body)
	if err != nil {
		log.Fatalln(err)
		//fmt.Fprintln(w1,err)
	}
	fmt.Fprintf(w1,string(content))
	data[cookiename] = new(tycontent)
	err = json.Unmarshal(content,data[cookiename])
	if err != nil {
		log.Println(err)
		//fmt.Fprintln(w1,err)
	}
	fmt.Fprintf(w1,"success")
}

func content(w2 http.ResponseWriter, r2 *http.Request)  {
	c := r2.Cookies()
	//fmt.Fprintf(w2,"查询中")
	//c,_ := r.Cookie(titles.Name)
	if len(c) == 0 {
		fmt.Fprintln(w2,"请登陆")
		fmt.Fprintln(w2,"跳转中...")
		time.Sleep(100 *time.Millisecond)
		http.Redirect(w2,r2,"http://localhost:9090/body",301)
	}
	var cookiename string
	for _,cookieq := range c {
		cookiename = cookieq.Value
	}
	if data[cookiename] == nil {
		w2.WriteHeader(403)
		return
	}
	fmt.Fprintf(w2,"您的信息：\n")
	con,_ := json.Marshal(data[cookiename])
	fmt.Fprintln(w2,string(con))
}
/*
func read()  {

}
*/
func test(ww http.ResponseWriter,rr *http.Request)  {
	fmt.Fprintf(ww,"success!!")
}

func main() {
//	http.HandleFunc("/content",content)
	http.HandleFunc("/content",content)
	http.HandleFunc("/m_content",m_content)
	http.HandleFunc("/test",test)
	http.HandleFunc("/body",main_body)
	err := http.ListenAndServe("localhost:9090",nil)
	if err != nil {
		log.Fatalln(err)
	}
}

