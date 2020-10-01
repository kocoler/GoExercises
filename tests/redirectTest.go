package main

import "github.com/gin-gonic/gin"

func main()  {
	var a *gin.Engine
	a = gin.Default()
	a.GET("/1", func(c *gin.Context) {
		c.Redirect(301,"http://baidu.com")
	})
	a.Run(":9999")
}
