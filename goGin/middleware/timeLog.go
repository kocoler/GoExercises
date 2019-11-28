package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func TimeLog() gin.HandlerFunc {
	return func(context *gin.Context) {
		startTime := time.Now()
		context.Next()
		endTime := time.Now()
		log.Println("Cost:",endTime.Sub(startTime).Nanoseconds())
		//fmt.Println(startTime)
	}
}
