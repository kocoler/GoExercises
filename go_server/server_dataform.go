// WebService project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)
var data = make(map[string]string,10)
var i int
func body(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "正在验证....\n")
	//time.Sleep(1 *time.Millisecond)
	//var data map[string]string
	name := r.FormValue("name")
	if len(data[name]) == 0 {
		i ++
		fmt.Fprintf(w,"第"+strconv.Itoa(i)+"次操作\n")
		fmt.Fprintf(w,"注册用户\n")
		password := r.FormValue("password")
		//fmt.Fprintln(w,password)
		data[name] = password
		fmt.Fprintf(w,"注册成功\n")
	} else {
		i ++
		fmt.Fprintf(w,"第"+strconv.Itoa(i)+"次操作\n")
		password := r.FormValue("password")
		if data[name] == password {
			fmt.Fprintf(w,"登入成功！\n")
		}else {
			fmt.Fprintf(w,"修改密码\n")
			data[name] = password
			fmt.Fprintf(w,"修改成功\n")
		}

	}
	/*a, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	r.Body.Close()
	//panic(a)
	fmt.Fprintf(w,string(a))*/


}


func main() {
	http.HandleFunc("/", body)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
