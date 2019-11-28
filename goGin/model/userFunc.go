package model

import (
	"goGin/funcs"
	"log"
)

type User struct {
	Id        int
	Name      string	`json:"username"`
	Password  string	`json:"-"`
	Salt	  string	`json:"-"`
	Sex       string	`json:"sex"`
	Favourite string	`json:"favourite"`
}

func CheckUser(username string) bool {
	var tmpUser  User
	if Db.Self.Model(&User{}).Where(User{Name:username}).First(&tmpUser); len(tmpUser.Name) != 0 {
		return true
	}
	return false
}

func LogInM(username string, password string) bool {
	var tmpUser User
	log.Println("登入用户："+username)
	Db.Self.Model(&User{}).Where(User{Name:username}).Find(&tmpUser)
	//log.Println(tmpUser)
	if  funcs.Check(password, tmpUser.Password, tmpUser.Salt) {
		return true
	}
	return false
}

func CreateUser(username string, password string) error {
	password, salt ,err := funcs.PasswordHash(password)
	if err != nil {
		return err
	}
	var tmpUser = &User{
		Name:      username,
		Password: password,
		Salt:      salt,
	}
	if err := Db.Self.Create(tmpUser).Error; err != nil {
		return err
	}
	return nil
}

func UpdatePassword(username string, password string) error {
	password, salt, err := funcs.PasswordHash(password)
	if err != nil {
		log.Println("111")
		return err
	}
	var tmpnUser  = &User {
		Name:       username,
		Password:  password,
		Salt:       salt,
	}
	if err := Db.Self.Model(&User{}).Where(&User{Name:username}).Update(&tmpnUser).Error; err != nil {
		return err
	}
	return nil
}

func UpdateInfo(tmpUser User) error {
	if err := Db.Self.Model(&User{}).Where(&User{Name:tmpUser.Name}).Update(&tmpUser).Error; err != nil {
		return err
	}
	return nil
}

func UserInfo(username string) (User,error) {
	var tmpUser User
	if err := Db.Self.Model(&User{}).Where(&User{Name:username}).Find(&tmpUser).Error; err != nil {
		return User{},err
	}
	return tmpUser,nil
}