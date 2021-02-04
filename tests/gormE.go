package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"sort"
	"time"
)

type users_info struct {
	Su       int       `gorm:"su"`
	School   string    `gorm:"school"`
	SendTime []uint8   `gorm:"send_time"`
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

	var tmp users_info
sort.Search()
	var a int
	db.Table("users_info").Count(&a)
	fmt.Println(a)
	var b int
	rand.Seed(time.Now().UnixNano())
	b = rand.Intn(a)
	fmt.Println(b)
/*
	for i := 0; i <= 3; i++ {
		b = rand.Intn(a)
		fmt.Println(b)
	}
*/
	/*for i := 0; i < 3; i ++ {
		if err := db.Order().Take(&tmp).Error; err != nil {
			log.Println(err)db.Order().Take(&tmp)
		}
		fmt.Println(tmp)
	}*/
	if err := db.Order("rand()").Take(&tmp).Error; err != nil {
		log.Println(err)
	}
	fmt.Println(tmp.Su)


	Db =


}
