package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goGin/handler/cookiesFuncs"
	"log"
)

func Welcome() gin.HandlerFunc {
	/*c.JSON(200,gin.H{
		"yeh" : "Welcome to my ginny!",
	})*/
	return func(context *gin.Context) {
		fmt.Println("welcome!!")
		name, err := cookiesFuncs.CheckCookiesName(context)
		if err != nil {
			log.Println(err)
			context.JSON(400, gin.H{})
		}
		context.JSON(200, gin.H{
			"welcome": name,
		})
		context.Next()
	}
}
