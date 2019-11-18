Learn more or give us feedback
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var data = make(map[string]string,10)
var i int

func body(w http.ResponseWriter, r *http.Request)  {
	//r.FormValue()
	defer r.Body.Close()
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var titles struct{
		Name string
		Password string
	}
	err = json.Unmarshal(content, &titles)
	if err != nil {
		log.Fatalln(err)
	}
	//fmt.Fprintf(w,"success")
	if len(data[titles.Name]) == 0 {
		i ++
		fmt.Fprintf(w,"第"+strconv.Itoa(i)+"次操作\n")
		fmt.Fprintf(w,"注册用户\n")
		password := titles.Password
		//fmt.Fprintln(w,password)
		data[titles.Name] = password
		fmt.Fprintf(w,"注册成功\n")
	} else {
		i ++
		fmt.Fprintf(w,"第"+strconv.Itoa(i)+"次操作\n")
		password := titles.Password
		if data[titles.Name] == password {
			fmt.Fprintf(w,"登入成功！\n")
		}else {
			fmt.Fprintf(w,"修改密码\n")
			data[titles.Name] = password
			fmt.Fprintf(w,"修改成功\n")
		}

	}
	//json.Unmarshal(content,&titles)
	//fmt.Fprintln(w,titles)
}

func main() {
	http.HandleFunc("/json",body)
	err := http.ListenAndServe(":9090",nil)
	if err != nil {
		log.Fatalln(err)
	}
}

