package main

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type a struct {
	Sid string `json:"sid"`
	Name string `json:"name"`
}

func main(){
	str := `{
	"sid": "2019214300",
		"name": "啦啦啦"
}`
	var result a
	err := json.Unmarshal([]byte(str),&result)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}