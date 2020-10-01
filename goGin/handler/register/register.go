package register

import (
	"github.com/gin-gonic/gin"
	"goGin/handler/cookiesFuncs"
	"goGin/model"
	"log"
)

type ReadUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	var userData ReadUser
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if model.CheckUser(userData.Username) {
		c.JSON(401, gin.H{
			"message": "User already existed!",
		})
		return
	}
	if err := model.CreateUser(userData.Username, userData.Password); err == nil {
		cookiesFuncs.SetCookies(c, userData.Username)
		c.JSON(200, gin.H{
			"message": "Success!!",
		})
	} else {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
	}

	return
}
