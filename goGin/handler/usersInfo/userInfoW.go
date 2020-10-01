package usersInfo

import (
	"github.com/gin-gonic/gin"
	"goGin/handler/cookiesFuncs"
	"goGin/model"
	"log"
)

/*type User struct {
	Name string
	Sex string
	Favourite string
}*/

func UsersInfoWrite(c *gin.Context) {
	var tmpUser model.User
	var err error
	tmpUser.Name, err = cookiesFuncs.CheckCookiesName(c)
	if err != nil {
		log.Println(err)
	}
	if err = c.BindJSON(&tmpUser); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bed Request",
		})
	}
	if err = model.UpdateInfo(tmpUser); err != nil {
		log.Println(err)
	}
	c.JSON(200, gin.H{
		"message": "Succeed!",
	})
}
