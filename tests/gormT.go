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
	SendTime []uint8   `gorm:"time_time"`
	TimeDate []uint8   `gorm:"time_date"`
	TimeForm int       `gorm:"time_form"`
	TimeEnd  int       `gorm:"time_end"`
}

func main(){
	db, err := gorm.Open("mysql","root:ccnudx@tcp(localhost:3306)/users")
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)
	//timeN := time.Now()
	//timeNow := timeN.Format("2006-01-02 15:04:05")

	var tmpUsers []users_info
/*
	tmp := users_info{
		//Su:3,
		School:"information",
		SendTime:[]uint8(timeNow),
		TimeDate:[]uint8("2020-01-12"),
		TimeForm:8,
		TimeEnd:12,
	}
	*/
	/*
	if err := db.Model(&users_info{}).Update(tmp).Error; err != nil {
		fmt.Println(err)
	}*/
	//fmt.Println(string(tmp.TimeDate))
	//fmt.Println(tmp)	timeNow := tNow.Format("2006-01-02 15:04:05")
	if err := db.Model(users_info{}).Where(users_info{TimeDate:[]uint8("2020-01-12")}).Find(&tmpUsers).Error; err != nil {
		log.Println(err)
	}
	for _,v := range tmpUsers {
		if v.TimeForm - 8 <= 2 || 8 - v.TimeForm <= 2 {
			if v.TimeEnd - 12 <= 2 || 12 - v.TimeEnd >= 2 {
				fmt.Println(v)
			}
		}
	}
}
