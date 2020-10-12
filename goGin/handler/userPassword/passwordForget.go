package userPassword

import (
	"github.com/gin-gonic/gin"
	"goGin/model"
	"log"
)

type ReadUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Forget(c *gin.Context) {
	var tmpUser ReadUser
	if err := c.BindJSON(&tmpUser); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	//log.Println("010101")
	if err := model.UpdatePassword(tmpUser.Username, tmpUser.Password); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}
	c.JSON(200, gin.H{
		"messege": "Succeed!",
	})
}
