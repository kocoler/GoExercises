package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type user struct {
	User_id int    `gorm:"user_id"`
	Qqq     int    `gorm:"qqq"`
	Qwq     string `grom:"qwq"`
}

func main() {
	db, err := gorm.Open("mysql","root:ccnudx@tcp(localhost:3306)/mini_project")
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)

	Db := db
	var tmp user

	//tmp.Sid = "1"
	tmp.Qqq = 2
	//tx := Db.Begin()
	//tx = tx.Model(users{}).Where(users{Nike_name:"1"})
	Db.Model(&user{}).Create(user{})
	//Db = Db.Model(users{}).Where(users{Gender:"3"})
	//tx = tx.Raw("`users`.`gender` = ? ",3)
	//Db.Raw("SELECT * FROM users ")
	//Db.Exec("gender = ")
	//b := 6
	//Db = Db.Where(" ABS(gender - ?) > 2",b)


}