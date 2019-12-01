package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var a string
	b := []string{
		"qqq",
	}
	a = "www"
	b = append(b,a)
	fmt.Println(b)
	for _,v := range b{
		fmt.Println(v)
	}
	type content struct {
		Content []string	//`json:"content"`
	}
	var c content
	c.Content = b
	d,_ := json.Marshal(b)
	fmt.Println(string(d))
}
