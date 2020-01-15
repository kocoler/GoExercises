package main


import (
	"fmt"
	"time"
	//"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

type users_info struct {
	Su       int       `gorm:"su"`
	School   string    `gorm:"school"`
	SendTime time.Time   `gorm:"send_time"`
	TimeDate string   `gorm:"time_date"`
	TimeForm int       `gorm:"time_form"`
	TimeEnd  int       `gorm:"time_end"`
}

func main() {
	db, err := gorm.Open("mysql", "root:ccnudx@tcp(localhost:3306)/users?parseTime=true")
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	db.SingularTable(true)

	var tmp []users_info

	db.Raw("SELECT * FROM users_info WHERE time_date & 11000000 > 10000000 ").Find(&tmp)
	//fmt.Println(db)

	fmt.Println(tmp)

}