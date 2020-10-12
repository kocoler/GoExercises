package main

import (
	"fmt"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type users struct {
	Sid       string       `gorm:"sid"`
	College   string    `gorm:"college"`
	NikeName string   `gorm:"nike_name"`
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
	var nikename []string
	Db.Model(&users{}).Where(users{College:"2"}).Pluck("nike_name",&nikename)
	fmt.Println(nikename)
	//Db = Db.Offset(2)
}