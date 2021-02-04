package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main(){
	res, err := http.Get("http://example.com/")
	if err!=nil{
		log.Println(err.Error())
		return
	}
	body,err:= ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	fmt.Println(string(body))

}