package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

const dns = "root:CCNU200181_dx@/ginUsers?charset=utf8&parseTime=True&loc=Local"

type Database struct {
	Self *gorm.DB
}

var Db *Database

func getDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func (db *Database) Init() {
	newDb, err := getDatabase()
	if err != nil {
		log.Print("数据库初始化错误")
		log.Println(err)
	}
	Db = &Database{Self: newDb}
	//return
}

func (db *Database) Close() {
	if err := Db.Self.Close(); err != nil {
		log.Print("数据库关闭错误")
		log.Println(err)
	}
	return
}
