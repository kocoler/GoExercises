package login

import (
	"github.com/gin-gonic/gin"
	"goGin/handler/cookiesFuncs"
	"goGin/model"
	"time"
)

type ReadUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var userData ReadUser
	if err := c.BindJSON(&userData); err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if model.LogInM(userData.Username, userData.Password) {
		cookiesFuncs.SetCookies(c, userData.Username)
		c.JSON(200, gin.H{
			"message": "Authentication Success!",
		})
		time.Sleep(100 * time.Microsecond)
	} else {
		c.JSON(401, gin.H{
			"message": "Authentication Failed.",
		})
		return
	}
	return
}
