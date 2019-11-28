package funcs

import (
	"github.com/gin-gonic/gin"
	"goGin/handler/cookiesFuncs"
	"log"
)

func Test1(c *gin.Context) {
	log.Println(cookiesFuncs.CheckCookiesName(c))
}
