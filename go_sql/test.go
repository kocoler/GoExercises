package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

const (
	DNS = "kocoler:CCNU200181_dx@/go_test?charset=utf8&parseTime=True&loc=Local"
)

var db  *gorm.DB

type User struct {
	//gorm.Model
	Id       int   `gorm:"primary_key"`//对应数据表的自增id
	Name string    `gorm:"not null"`
	Age int
	//Updated_at
}

func main()  {
	db, err := gorm.Open("mysql", DNS)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	db.SingularTable(true)
	//var a User
	user := &User {
		Id : 101,
		Name: "42",
		Age: 11,
	}
	//db.Where("id=?",8).Delete(&User{})
	db.Save(&user)
	//db.Model(&user).Update("CreatAt",time.Now())
	//var b  = User{}
	fmt.Println(time.Now())
}
