package gorm

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type ServiceUser struct {
	ID        int       `json:"id" gorm:"primaryKey;column:id;type:int auto_increment;comment:id"` // id
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:timestamp"`
	DeletedAt time.Time `json:"deleted_at" gorm:"column:deleted_at;type:timestamp"`
	UserID    int       `json:"user_id" gorm:"column:user_id;type:int;not null;default:0;comment:用户 ID"`            // 用户 ID
	ServiceID int       `json:"service_id" gorm:"column:service_id;type:int;not null;default:0;comment:服务 ID"`      // 服务 ID
	Type      int32     `json:"type" gorm:"column:type;type:tinyint;not null;default:0;comment:用户身份：0: 管理员 1: 开发者"` // 用户身份：0: 管理员 1: 开发者
}

// TableName returns the table name of the ServiceUser model
func (s *ServiceUser) TableName() string {
	return "service_user"
}

// Create creates a data of the ServiceUser model.
func (s *ServiceUser) Create() error {
	return DB.Create(s).Error
}

// Save saves a data of the ServiceUser model.
func (s *ServiceUser) Save() error {
	return DB.Save(s).Error
}

// Update updates a data of the ServiceUser model.
func (s *ServiceUser) Update() error {
	return DB.Update(s).Error
}

// Delete deletes a data of the ServiceUser model.
func (s *ServiceUser) Delete() error {
	return DB.Delete(s).Error
}

// GetServiceDevelopers returns the developers of the service.
func GetServiceDevelopers(serviceID int) ([]string, error) {
	var res []string
	err := DB.Model(&ServiceUser{}).
		Where("service_id = ?", serviceID).
		Where("type = ?", 1).
		Joins("INNER JOIN user on service_user.user_id = user.id").
		Pluck("user.nickname", &res).Error

	return res, err
}

func main() {
	DB, err := gorm.Open("mysql", "root:ccnudx@tcp(localhost:3306)/mae")
	defer DB.Close()
	if err != nil {
		log.Println(err)
	}
	DB.SingularTable(true)
}
