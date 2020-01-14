package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type users_info struct {
	Su       int       `gorm:"su"`
	School   string    `gorm:"school"`
	SendTime []uint8   `gorm:"send_time"`
	TimeDate []uint8   `gorm:"time_date"`
	TimeForm int       `gorm:"time_form"`
	TimeEnd  int       `gorm:"time_end"`
}

func main() {
	db, err := gorm.Open("mysql","root:ccnudx@tcp(localhost:3306)/users")
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)

	Db := db
	var tmp []users_info
	Db = Db.Model(users_info{}).Where(users_info{TimeForm:6}).Find(&tmp)
	//fmt.Println(tmp)
	Db = Db.Model(users_info{}).Where(users_info{TimeEnd:11}).Find(&tmp)
	fmt.Println(tmp)

	for i := 0; i <= n; i ++ {
		Db.Model().Where(users_info{})
	}
Db.

}