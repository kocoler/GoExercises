package main

import "fmt"

type User interface {
	Func1()
}

func (s *Users) Func1() {
	fmt.Println("func1")
}
type Users struct {
	SuInfo string
	User
}
func (s *Users) UserInfo()  {
	fmt.Println(s.SuInfo)
}
func (s *Users)UserFunc() {
	fmt.Println("qqq")
}
func main() {
	var d Users
	d.SuInfo = "1"
	d.Func1()
}
