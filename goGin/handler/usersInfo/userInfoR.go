package usersInfo

import (
	"github.com/gin-gonic/gin"
	"goGin/handler/cookiesFuncs"
	"goGin/model"
	"log"
)

func UserInfoR(c *gin.Context)  {
	tmpUsername := c.Query("name")
	username, err := cookiesFuncs.CheckCookiesName(c)
	log.Println(username)
	if err != nil {
		c.JSON(400,gin.H{

		})
		return
	}
	if tmpUsername != username {
		c.JSON(403,gin.H{
			"message":"Forbidden",
		})
		return
	}
	content, err := model.UserInfo(username)
	if err != nil {
		c.JSON(400,gin.H{

		})
		return
	}
	c.JSON(200,content)
}
