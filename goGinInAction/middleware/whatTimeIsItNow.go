package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"sync"
	"time"
)

var I int = 1
var lock sync.Mutex

func Time() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("正在操作中......")
		log.Println("服务启动后的第" + strconv.Itoa(I) + "次操作")
		lock.Lock()
		I++
		lock.Unlock()
		log.Println(time.Now())
		//context.JSON(200,time.Now())
		context.Next()
	}
}
