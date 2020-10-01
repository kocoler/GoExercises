package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	DNS = "root:CCNU200181_dx@/ginUsers?charset=utf8&parseTime=True&loc=Local"
)

var db *gorm.DB

type user struct {
	//gorm.Model
	Id   int    //`gorm:"primary_key"`//对应数据表的自增id
	Name string //`gorm:"not null"`
	Sex  string
	//Updated_at time.Time
}

func main() {
	db, err := gorm.Open("mysql", DNS)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	//db.SingularTable(true)
	//var a User
	users := new(user)
	users.Id = 111
	//users.Name = "41"
	//db.AutoMigrate(&user{})
	//db.Where("id=?",8).Delete(&User{})
	//db.Create(users)
	//db.Model(&user).Update("CreatAt",time.Now())
	//var b  = User{}
	/*	db.Model(&users).Update(user{
		Id:   1,
		Name: "111",
		Sex:"11",
	})*/

	/*var a  = &user{
		//Id:   0,
		Name: "42",
		Sex:  "111",
	}*/
	var b user
	db.Model(&user{}).Where(&user{Name: "2"}).Find(&b)
	fmt.Println(b.Name)

	//fmt.Println(time.Now())
}
