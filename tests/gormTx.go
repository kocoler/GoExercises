package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type users struct {
	Sid       string       `gorm:"sid"`
	College   string    `gorm:"college"`
	Nike_name string   `gorm:"nike_name"`
	Gender string   `gorm:"gender"`
	Grade string       `gorm:"grade"`
}

func main() {
	db, err := gorm.Open("mysql","root:ccnudx@tcp(localhost:3306)/mini_project")
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)

	Db := db
	var tmp []users

	tx := Db.Begin()
	tx = tx.Model(users{}).Where(users{Nike_name:"1"})
	//Db = Db.Model(users{}).Where(users{Nike_name:"1"})
	tx = tx.Model(users{}).Where(users{College:"2"})
	//Db = Db.Model(users{}).Where(users{College:"2"})
	//Db = Db.Model(users{}).Where(users{Gender:"3"})
	//tx = tx.Raw("`users`.`gender` = ? ",3)
	//Db.Raw("SELECT * FROM users ")
	//Db.Exec("gender = ")

	tx.Find(&tmp)

	//Db = Db.Offset(2)

	fmt.Println(tmp)

}