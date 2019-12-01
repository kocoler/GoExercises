package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	Id        int
	Name      string	`json:"username"`
	Password  string	`json:"-"`
	Salt	  string	`json:"-"`
	Sex       string	`json:"sex"`
	Favourite string	`json:"favourite"`
}

func main() {
	db, err := gorm.Open("mysql","root:CCNU200181_dx@/ginUsers?charset=utf8&parseTime=True&loc=Local")
	if err !=nil {
		log.Println("shujuku",err)
	}
	var a []User
	db.Model(&User{}).Where(User{Name:"1317"}).Find(&a)
	fmt.Println(a)
}
