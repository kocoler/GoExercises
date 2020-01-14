package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kocoler/GoExercises/miniProject/middleware"
	"github.com/kocoler/GoExercises/miniProject/model"
	"math/rand"
	"time"
)

type UserInfo struct {
	Sid string `json:"-" gorm:"sid"`
	NickName string `json:"nick_name" gorm:"nick_name"`
	College string `json:"college" gorm:"college"`
	Gender string `json:"gender" gorm:"gender"`
	PersonalizedSignature string `json:"personalized_signature" gorm:"personalized_signature"`
	ContactWay string `json:"contact_way" gorm:"contact_way"`
}

type LoginInfo struct {
	Sid string `json:"sid"`
	Pwd string `json:"pwd"`
}

func UserLogin(c *gin.Context) {
	//a, _ := c.Get("token2")
	var tmpUser model.SuInfo
	var tmpLoginInfo LoginInfo
	if err := c.BindJSON(&tmpLoginInfo); err != nil {
		c.JSON(400,gin.H{
			"message":"Bad Request!",
		})
		return
	}
	tmpUser, err := model.GetUserInfoFormOne(tmpLoginInfo.Sid, tmpLoginInfo.Pwd)
	if  err != nil {
		c.JSON(400,gin.H{
			"message":err,
		})
		return
	}
	fmt.Print(tmpUser)

	user := UserInfo{
		Sid:                   tmpUser.User.Usernumber,
		NickName:              getRandomString(8),
		College:               tmpUser.User.DeptName,
		Gender:                getGender(tmpUser.User.Xb),
		PersonalizedSignature: "",
		ContactWay:            "",
	}

	//写入用户数据到数据库

	fmt.Println(user)

	c.Header("token",middleware.ProduceToken(user.Sid))
}

func Test(c *gin.Context) {
	a,b := c.Get("uid")
	fmt.Println(a,b)
}

func getGender(n string) string {
	if n == "1" {
		return "男"
	}else {
		return "女"
	}
}

func  getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
