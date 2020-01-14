package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

const dns = "root:ccnudx@tcp(localhost:3306)/mini_project"

type Database struct {
	Self *gorm.DB
}

var Db *Database

func getDatabase() (*gorm.DB, error) {
	db, err := gorm.Open("mysql",dns)
	if err != nil {
		fmt.Print("getDatabase")
		log.Println(err)
	}
	return db, err
}

func (db *Database) Init() {
	newDb, err := getDatabase()
	if err != nil {

	}
	Db = &Database{Self:newDb}
}

func (db *Database) Close() {
	if err := Db.Self.Close(); err != nil {

	}
	return
}