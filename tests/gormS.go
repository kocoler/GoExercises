package main


import (
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

	var tmp []users_info

	if err := db.Model(users_info{}).Where(users_info{School:"information",TimeForm:6,})

}
