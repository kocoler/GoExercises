package model

import (
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

const dns  = "root:CCNU200181_dx@/ginUsers?charset=utf8&parseTime=True&loc=Local"

type Database struct {
	Self *gorm.DB
}

var Db *Database

func getDatabase() (*gorm.DB,error) {
	db, err := gorm.Open("mysql",dns)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	return db, nil
}

func (db *Database) Init() {
	newdb, err := getDatabase()
	if err != nil {
		log.Fatal(err)
	}
	Db = &Database{
		Self:newdb,
	}
}

func (db *Database) Close() {
	if err := Db.Self.Close(); err!=nil {
		log.Println(err)
	}
}

